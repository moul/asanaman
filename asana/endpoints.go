package asana

import "context"

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
