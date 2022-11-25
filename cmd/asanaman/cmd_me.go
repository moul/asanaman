package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"
	"moul.io/u"
)

func doMe(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	g.rootLogger.Debug("me", zap.Strings("args", args), zap.Any("g", g))

	ret, err := g.client.Me(ctx, nil)
	if err != nil {
		return fmt.Errorf("me: %w", err)
	}

	fmt.Println(u.PrettyJSON(ret))
	return nil
}
