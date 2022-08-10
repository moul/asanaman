package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
	g.rootLogger.Debug("workspaces", zap.Int("count", len(*ret)))

	if g.JSONFormat {
		fmt.Println(u.PrettyJSON(ret))
		return nil
	}

	// ascii rendering
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"GID", "NAME", "RESOURCE TYPE"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(false)
	for _, entry := range *ret {
		gid := entry.GID
		name := entry.Name
		resourceType := entry.ResourceType
		table.Append([]string{gid, name, resourceType})
	}
	total := fmt.Sprintf("%d", len(*ret))
	table.SetFooter([]string{"Total", "", total})
	table.Render()

	return nil
}
