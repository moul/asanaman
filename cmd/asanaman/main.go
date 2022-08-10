package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/peterbourgon/ff/v3"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"moul.io/asanaman/asana"
	"moul.io/climan"
	"moul.io/srand"
	"moul.io/zapconfig"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		if !errors.Is(err, flag.ErrHelp) {
			fmt.Fprintf(os.Stderr, "error: %+v\n", err)
		}
		os.Exit(1)
	}
}

// nolint:maligned
var g struct {
	// opts
	Debug     bool
	Token     string `json:"-"` // sensitive
	Domain    string
	CachePath string

	// internal
	rootLogger *zap.Logger
	client     *asana.Client

	// subcommand opts
	FilterWorkspace string `json:"workspace,omitempty"`
	FilterTeam      string `json:"team,omitempty"`
	FilterArchived  bool   `json:"archived,omitempty"`
}

func run(args []string) error {
	// defaults
	g.CachePath = filepath.Join(".", ".asanaman-cache")

	commonFlags := func(fs *flag.FlagSet) {
		fs.BoolVar(&g.Debug, "debug", g.Debug, "debug mode")
		fs.StringVar(&g.Token, "token", g.Token, "Asana token")
		fs.StringVar(&g.Domain, "domain", g.Domain, "Asana workspace")
		fs.StringVar(&g.CachePath, "cache-path", g.CachePath, "cache path")
	}

	// parse CLI
	root := &climan.Command{
		Name:           "asanaman",
		ShortUsage:     "asanaman [global flags] <subcommand> [flags] [args]",
		ShortHelp:      "More info on https://moul.io/asanaman.",
		FlagSetBuilder: commonFlags,
		FFOptions:      []ff.Option{ff.WithEnvVarPrefix("asanaman")},
		Subcommands: []*climan.Command{
			{Name: "me", Exec: doMe, FlagSetBuilder: func(fs *flag.FlagSet) { commonFlags(fs) }},
			{Name: "project-list", Exec: doProjectList, FlagSetBuilder: func(fs *flag.FlagSet) {
				commonFlags(fs)
				fs.StringVar(&g.FilterWorkspace, "filter-workspace", g.FilterWorkspace, "filter by workspace")
				fs.StringVar(&g.FilterTeam, "filter-team", g.FilterTeam, "filter by team")
				fs.BoolVar(&g.FilterArchived, "filter-archived", g.FilterArchived, "filter by archive status")
			}},
		},
	}
	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	if g.Token == "" {
		return fmt.Errorf("missing asana token (see https://app.asana.com/0/my-apps)") // nolint:goerr113
	}
	if g.Domain == "" {
		return fmt.Errorf("missing asana domain/workspace") // nolint:goerr113
	}

	// init runtime
	{
		// prng
		rand.Seed(srand.Fast())

		// concurrency
		// runtime.GOMAXPROCS(1)

		// logger
		config := zapconfig.New().SetPreset("light-console")
		if g.Debug {
			config = config.SetLevel(zapcore.DebugLevel)
		} else {
			config = config.SetLevel(zapcore.InfoLevel)
		}
		var err error
		g.rootLogger, err = config.Build()
		if err != nil {
			return fmt.Errorf("logger init: %w", err)
		}

		// asana
		{
			logger := g.rootLogger.Named("client")
			g.client, err = asana.New(g.Token, g.Domain, g.CachePath, logger)
			if err != nil {
				return fmt.Errorf("asana client: %w", err)
			}
		}
	}

	// run
	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
