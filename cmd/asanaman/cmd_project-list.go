package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
	g.rootLogger.Debug("projects", zap.Int("count", len(*ret)))

	if g.JsonFormat {
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
