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

	opts.rootLogger.Debug("init", zap.Strings("args", args), zap.Any("opts", opts))
	fmt.Println(opts.client)

	ret, err := opts.client.Me(ctx)
	if err != nil {
		return fmt.Errorf("me: %w", err)
	}

	fmt.Println(u.PrettyJSON(ret))
	return nil
}
