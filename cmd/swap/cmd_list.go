package main

import (
	"context"

	"github.com/urfave/cli/v3"
)

var subCmd = &cli.Command{
	Name:   "list",
	Usage:  "display convertable pairs",
	Action: subAction,
	Hidden: true,
}

var subAction = func(c context.Context, cmd *cli.Command) error { return nil }
