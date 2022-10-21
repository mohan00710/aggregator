package commands

import (
	"aggregator/config"
	"aggregator/loadbalance"
	"aggregator/server"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func RunCommand() *cli.Command {
	return &cli.Command{
		Name:    "run",
		Aliases: []string{"start"},
		Flags:   append([]cli.Flag{}, InitCommand().Flags...),
		Before: func(cli *cli.Context) error {
			err := runCommand(cli, "init")
			if err != nil {
				return err
			}

			config.Load()
			loadbalance.LoadFromConfig()
			return nil
		},
		Action: func(context *cli.Context) error {
			wg := errgroup.Group{}
			wg.Go(func() error {
				return server.NewServer()
			})

			wg.Go(func() error {

				return server.NewManageServer()
			})

			return wg.Wait()
		},
		Subcommands: []*cli.Command{},
	}

}
