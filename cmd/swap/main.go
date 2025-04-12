package main

import (
	"context"
	"fmt"
	"os"

	markdown "github.com/MichaelMure/go-term-markdown"
	"github.com/lfaoro/swap/pkg/types"
	"github.com/urfave/cli/v3"
)

var BuildVersion string = "v0.0.1-dev"
var BuildDate string = "unset"
var BuildSHA string = "unset"
var APIURL string = "api.swapcli.com:443"

func main() {
	appcmd := &cli.Command{
		Authors: []any{
			map[string]string{
				"name":  "Leonardo Faoro",
				"email": "swap@leonardofaoro.com",
			},
		},
		Name:                   "swap",
		EnableShellCompletion:  true,
		UseShortOptionHandling: true,
		Suggest:                true,

		Version: fmt.Sprintf("Version %s\nBuild date: %s\nBuild SHA: %s", BuildVersion, BuildDate, BuildSHA),
		ExtraInfo: func() map[string]string {
			return map[string]string{
				"Build version": BuildVersion,
				"Build date":    BuildDate,
				"Build sha":     BuildSHA,
			}
		},

		Usage:     "Crypto Swaps Terminal",
		UsageText: `Swap is a Terminal-UI that facilitates secure cross-chain asset swaps.`,

		Before: func(c context.Context, cmd *cli.Command) (context.Context, error) {
			ctx := context.WithValue(c, types.APIURLKey, APIURL)
			return ctx, nil
		},
		Action: mainCmd,

		Commands: []*cli.Command{
			subCmd,
		},

		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "no-logs",
				Aliases:     []string{"n"},
				DefaultText: "do not store transaction logs",
				Value:       false,
				Action: func(context.Context, *cli.Command, bool) error {
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "legal",
				Aliases: []string{"terms"},
				Usage:   "print legal disclaimer",
				Action: func(context.Context, *cli.Command, bool) error {
					res := markdown.Render(legalDisclaimer, 80, 6)
					fmt.Println(string(res))
					os.Exit(0)
					return nil
				},
			},
			&cli.BoolFlag{
				Name:    "debug",
				Aliases: []string{"d"},
				Usage:   "enable debug mode with verbose logging",
				Value:   false,
			},
		},
	}

	err := appcmd.Run(context.Background(), os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
