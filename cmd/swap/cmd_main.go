// SPDX-License-Identifier: AGPL-3.0-only
// Copyright (c) 2024 Leonardo Faoro. All rights reserved.
// Use of this source code is governed by the AGPL-3.0 license
// found in the LICENSE file.
package main

import (
	"context"
	"fmt"
	"os"

	"crypto/tls"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/lfaoro/swap/app"
	pb "github.com/lfaoro/swap/gen/go/swap/v1"
	"github.com/lfaoro/swap/pkg/types"
	"github.com/urfave/cli/v3"
	"golang.org/x/term"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

var mainCmd = func(c context.Context, cmd *cli.Command) error {
	APIURL := c.Value(types.APIURLKey).(string)

	debug := cmd.Bool("debug")
	if debug {
		for k, v := range cmd.ExtraInfo() {
			fmt.Println(k, v)
		}
		fmt.Println("Build API", APIURL)
	}

	if !term.IsTerminal(int(os.Stdin.Fd())) {
		return fmt.Errorf("not an interactive terminal :(")
	}

	var opts []grpc.DialOption
	if APIURL == "localhost:8080" {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	} else {
		config := &tls.Config{
			InsecureSkipVerify: false,
			MinVersion:         tls.VersionTLS13,
		}
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(config)))
	}

	conn, err := grpc.NewClient(APIURL, opts...)
	if err != nil {
		return fmt.Errorf("failed to create connection: %w", err)
	}
	defer conn.Close()
	client := pb.NewCoinServiceClient(conn)

	cfg, err := app.NewConfig()
	if err != nil {
		return fmt.Errorf("failed to create config: %w", err)
	}

	m := app.NewTSwapUI(cfg, client, debug)
	_, err = tea.NewProgram(m, tea.WithOutput(os.Stderr)).
		Run()

	if err != nil {
		return fmt.Errorf("failed to run app: %w", err)
	}

	return nil
}
