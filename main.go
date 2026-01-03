package main

import (
	"chihqiang/depctl/cmdx"
	"chihqiang/depctl/flagx"
	"context"
	"os"

	"github.com/chihqiang/logx"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:  "depctl",
		Usage: "Push it, roll it, own it",
		Flags: flagx.SSHFlags(),
		Commands: []*cli.Command{
			cmdx.Publish(),
			cmdx.History(),
			cmdx.Rollback(),
		},
	}
	if err := app.Run(context.Background(), os.Args); err != nil {
		logx.Error("%v", err)
		os.Exit(1)
	}
}
