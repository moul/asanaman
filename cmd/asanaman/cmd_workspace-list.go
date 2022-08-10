package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"
	"moul.io/u"
)

func doWorkspaceList(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	g.rootLogger.Debug("workspace-list", zap.Strings("args", args), zap.Any("g", g))

	ret, err := g.client.WorkspaceList(ctx)
	if err != nil {
		return fmt.Errorf("workspace-list: %w", err)
	}

	fmt.Println(u.PrettyJSON(ret))
	g.rootLogger.Debug("workspaces", zap.Int("count", len(*ret)))
	return nil
}
