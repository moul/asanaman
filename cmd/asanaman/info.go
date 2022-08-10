package main

import (
	"context"
	"flag"
	"fmt"

	"go.uber.org/zap"
)

func doInfo(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return flag.ErrHelp
	}

	opts.rootLogger.Debug("init", zap.Strings("args", args), zap.Any("opts", opts))
	fmt.Println(opts.client)
	return nil
}
