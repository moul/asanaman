package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"

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

var opts struct {
	Debug      bool
	Token      string `json:"-"` // sensitive
	Domain     string
	rootLogger *zap.Logger
	client     *asana.Client
}

func run(args []string) error {
	commonFlags := func(fs *flag.FlagSet) {
		fs.BoolVar(&opts.Debug, "debug", opts.Debug, "debug mode")
		fs.StringVar(&opts.Token, "token", opts.Token, "Asana token")
		fs.StringVar(&opts.Domain, "domain", opts.Domain, "Asana workspace")
	}

	// parse CLI
	root := &climan.Command{
		Name:           "asanaman",
		ShortUsage:     "asanaman [global flags] <subcommand> [flags] [args]",
		ShortHelp:      "More info on https://moul.io/asanaman.",
		FlagSetBuilder: commonFlags,
		FFOptions:      []ff.Option{ff.WithEnvVarPrefix("asanaman")},
		Subcommands: []*climan.Command{
			{Name: "info", Exec: doInfo, FlagSetBuilder: func(fs *flag.FlagSet) { commonFlags(fs) }},
		},
	}
	if err := root.Parse(args); err != nil {
		return fmt.Errorf("parse error: %w", err)
	}

	if opts.Token == "" {
		return fmt.Errorf("missing asana token (see https://app.asana.com/0/my-apps)")
	}
	if opts.Domain == "" {
		return fmt.Errorf("missing asana domain/workspace")
	}

	// init runtime
	{
		// prng
		rand.Seed(srand.Fast())

		// concurrency
		// runtime.GOMAXPROCS(1)

		// logger
		config := zapconfig.New().SetPreset("light-console")
		if opts.Debug {
			config = config.SetLevel(zapcore.DebugLevel)
		} else {
			config = config.SetLevel(zapcore.InfoLevel)
		}
		var err error
		opts.rootLogger, err = config.Build()
		if err != nil {
			return fmt.Errorf("logger init: %w", err)
		}

		// asana
		opts.client, err = asana.New(opts.Token, opts.Domain)
		if err != nil {
			return fmt.Errorf("asana client: %w", err)
		}
	}

	// run
	if err := root.Run(context.Background()); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}
