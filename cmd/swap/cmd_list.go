package main

import (
	"context"

	"github.com/urfave/cli/v3"
)

var listCmd = &cli.Command{
	Name:   "list",
	Usage:  "display convertable pairs",
	Action: listFunc,
	Hidden: true,
}

var listFunc = func(c context.Context, cmd *cli.Command) error { return nil }
