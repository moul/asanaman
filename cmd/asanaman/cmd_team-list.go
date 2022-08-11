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

func doTeamList(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	g.rootLogger.Debug("team-list", zap.Strings("args", args), zap.Any("g", g))

	opts := asana.TeamListOpts{
		Workspace:    g.FilterWorkspace,
		User:         g.FilterUser,
		Organization: g.FilterOrganization,
	}
	ret, err := g.client.TeamList(ctx, opts)
	if err != nil {
		return fmt.Errorf("team-list: %w", err)
	}
	g.rootLogger.Debug("teams", zap.Int("count", len(*ret)))

	if g.JSONFormat {
		fmt.Println(u.PrettyJSON(ret))
		return nil
	}

	// ascii rendering
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"GID", "NAME", "RESOURCE TYPE", "DESCRIPTION", "ORGANIZATION", "VISIBILITY"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(false)
	for _, entry := range *ret {
		var (
			gid          = entry.GID
			name         = entry.Name
			resourceType = entry.ResourceType
			description  = entry.Description
			organization = entry.Workspace.Name
			// permalink    = entry.PermalinkURL
			visibility = entry.Visibility
		)
		if description == "" {
			description = "n/a"
		}

		table.Append([]string{gid, name, resourceType, description, organization, visibility})
	}
	total := fmt.Sprintf("Total: %d", len(*ret))
	table.SetFooter([]string{total, "", "", "", "", "", ""})
	table.Render()

	return nil
}
