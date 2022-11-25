package asana

import (
	"context"
	"fmt"
)

type MeOpts struct {
	OptFields string `url:"opt_fields,omitempty"`
}

func (c *Client) Me(ctx context.Context, opts *MeOpts) (*User, error) {
	if opts == nil {
		opts = &MeOpts{}
	}

	var user User
	return &user, c.Request(ctx, ReqOpts{Path: "users/me", Opts: *opts}, &user)
}

type ProjectListOpts struct {
	Workspace string `url:"workspace,omitempty"`
	Team      string `url:"team,omitempty"`
	Archived  bool   `url:"archived,omitempty"`
	OptFields string `url:"opt_fields,omitempty"`
}

func (c *Client) ProjectList(ctx context.Context, opts *ProjectListOpts) (*Projects, error) {
	if opts == nil {
		opts = &ProjectListOpts{}
	}

	if opts.OptFields == "" {
		opts.OptFields = "name,resource_type,archived,color,created_at,current_status,current_status_update,notes,public,start_on,workspace,completed,completed_at,completed_by,created_from_template,custom_fields,followers,icon,project_brief,team,workspace.name,team.name"
	}

	var projects Projects
	return &projects, c.Request(ctx, ReqOpts{Path: "projects", Opts: opts}, &projects)
}

type WorkspaceListOpts struct {
	OptFields string `url:"opt_fields,omitempty"`
}

func (c *Client) WorkspaceList(ctx context.Context, opts *WorkspaceListOpts) (*Workspaces, error) {
	if opts == nil {
		opts = &WorkspaceListOpts{}
	}

	if opts.OptFields == "" {
		opts.OptFields = "name,resource_type,email_domains,is_organization"
	}

	var workspaces Workspaces
	return &workspaces, c.Request(ctx, ReqOpts{Path: "workspaces", Opts: *opts}, &workspaces)
}

type TeamListOpts struct {
	Workspace    string `url:"workspace,omitempty"`
	User         string `url:"user,omitempty"`
	Organization string `url:"organization,omitempty"`
	OptFields    string `url:"opt_fields,omitempty"`
}

func (c *Client) TeamList(ctx context.Context, opts *TeamListOpts) (*Teams, error) {
	if opts == nil {
		opts = &TeamListOpts{}
	}

	var (
		getTeamsForUser      = opts.User != "" && opts.Organization != "" && opts.Workspace == ""
		getTeamsForWorkspace = opts.User == "" && opts.Organization == "" && opts.Workspace != ""
	)
	switch {
	case getTeamsForUser:
	case getTeamsForWorkspace:
	default:
		return nil, fmt.Errorf("invalid opts: should have Workspace OR User+Organization") // nolint:goerr113
	}

	if opts.OptFields == "" {
		opts.OptFields = "name,resource_type,description,html_description,organization,permalink_url,visibility,organization.resource_type,organization.name"
	}

	var teams Teams
	return &teams, c.Request(ctx, ReqOpts{Path: "teams", Opts: *opts}, &teams)
}
