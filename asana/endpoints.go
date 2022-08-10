package asana

import "context"

func (c *Client) Me(ctx context.Context) (*User, error) {
	var user User
	return &user, c.Request(ctx, ReqOpts{Path: "users/me"}, &user)
}

type ProjectListOpts struct{}

func (c *Client) ProjectList(ctx context.Context, opts ProjectListOpts) (*Projects, error) {
	var projects Projects
	return &projects, c.Request(ctx, ReqOpts{Path: "projects"}, &projects)
}
