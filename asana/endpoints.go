package asana

import "context"

func (c *Client) Me(ctx context.Context) (*User, error) {
	var user User
	return &user, c.Request(ctx, ReqOpts{Path: "users/me"}, &user)
}
