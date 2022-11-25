package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
	"moul.io/u"
)

func doWorkspaceList(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	g.rootLogger.Debug("workspace-list", zap.Strings("args", args), zap.Any("g", g))

	ret, err := g.client.WorkspaceList(ctx, nil)
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
	table.SetHeader([]string{"GID", "NAME", "RESOURCE TYPE", "EMAIL DOMAINS", "IS ORGANIZATION"})
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(false)
	for _, entry := range *ret {
		gid := entry.GID
		name := entry.Name
		resourceType := entry.ResourceType
		emailDomains := strings.Join(entry.EmailDomains, ",")
		isOrganization := fmt.Sprintf("%v", entry.IsOrganization)
		table.Append([]string{gid, name, resourceType, emailDomains, isOrganization})
	}
	total := fmt.Sprintf("Total: %d", len(*ret))
	table.SetFooter([]string{total, "", "", "", ""})
	table.Render()

	return nil
}
