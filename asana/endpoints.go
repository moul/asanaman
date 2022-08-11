package asana

import (
	"context"
	"fmt"
)

func (c *Client) Me(ctx context.Context) (*User, error) {
	var user User
	return &user, c.Request(ctx, ReqOpts{Path: "users/me"}, &user)
}

type ProjectListOpts struct {
	Workspace string `url:"workspace,omitempty"`
	Team      string `url:"team,omitempty"`
	Archived  bool   `url:"archived,omitempty"`
}

func (c *Client) ProjectList(ctx context.Context, opts ProjectListOpts) (*Projects, error) {
	var projects Projects
	return &projects, c.Request(ctx, ReqOpts{Path: "projects", Opts: opts}, &projects)
}

func (c *Client) WorkspaceList(ctx context.Context) (*Workspaces, error) {
	var workspaces Workspaces
	return &workspaces, c.Request(ctx, ReqOpts{Path: "workspaces"}, &workspaces)
}

type TeamListOpts struct {
	Workspace    string `url:"workspace,omitempty"`
	User         string `url:"user,omitempty"`
	Organization string `url:"organization,omitempty"`
	OptFields    string `url:"opt_fields,omitempty"`
}

func (c *Client) TeamList(ctx context.Context, opts TeamListOpts) (*Teams, error) {
	var (
		getTeamsForUser      = opts.User != "" && opts.Organization != "" && opts.Workspace == ""
		getTeamsForWorkspace = opts.User == "" && opts.Organization == "" && opts.Workspace != ""
	)

	if opts.OptFields == "" {
		opts.OptFields = "name,resource_type,description,html_description,organization,permalink_url,visibility,organization.resource_type,organization.name"
	}

	switch {
	case getTeamsForUser:
	case getTeamsForWorkspace:
	default:
		return nil, fmt.Errorf("invalid opts: should have Workspace OR User+Organization") // nolint:goerr113
	}

	var teams Teams
	return &teams, c.Request(ctx, ReqOpts{Path: "teams", Opts: opts}, &teams)
}
