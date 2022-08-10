package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"
	"moul.io/asanaman/asana"
	"moul.io/u"
)

func doProjectList(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	g.rootLogger.Debug("project-list", zap.Strings("args", args), zap.Any("g", g))

	opts := asana.ProjectListOpts{
		Workspace: g.FilterWorkspace,
		Team:      g.FilterTeam,
		Archived:  g.FilterArchived,
	}
	ret, err := g.client.ProjectList(ctx, opts)
	if err != nil {
		return fmt.Errorf("project-list: %w", err)
	}

	fmt.Println(u.PrettyJSON(ret))
	g.rootLogger.Debug("projects", zap.Int("count", len(*ret)))
	return nil
}
