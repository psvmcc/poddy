package main

import (
	"fmt"
	"log"
	"os"
	"poddy/pkg/commands"
	"poddy/pkg/types"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "poddy"
	app.Usage = "podman container manager"
	app.Version = fmt.Sprintf("%s (%s)", types.Version, types.Commit)
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "poddyendpoint",
			Usage:   "poddy server endpoint",
			EnvVars: []string{"PODDYENDPOINT"},
		},
		&cli.StringFlag{
			Name:    "poddytoken",
			Usage:   "poddy server auth token",
			EnvVars: []string{"PODDYTOKEN"},
		},
		&cli.StringFlag{
			Name:    "namespace",
			Aliases: []string{"n"},
			Usage:   "Namespace",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:    "info",
			Aliases: []string{"i"},
			Usage:   "Poddy server information",
			Action:  commands.Info,
		},
		{
			Name:    "server",
			Aliases: []string{"s"},
			Usage:   "Starts poddy server",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "verbose",
					Usage:   "Verbose logging",
					EnvVars: []string{"PODDY_VERBOSE"},
				},
				&cli.StringFlag{
					Name:    "bind",
					Usage:   "Bind address",
					Value:   "127.0.0.1:5922",
					EnvVars: []string{"PODDY_BIND"},
				},
				&cli.StringFlag{
					Name:    "self-exporter-bind",
					Usage:   "Metrics sefl exporter bind address",
					Value:   "127.0.0.1:9922",
					EnvVars: []string{"PODDY_SELF_EXPORTER_BIND"},
				},
				&cli.StringFlag{
					Name:    "config",
					Usage:   "Server config path",
					Value:   "config.yaml",
					EnvVars: []string{"PODDY_SERVER_CONFIG"},
				},
			},
			Action: commands.StartServer,
		},
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "Get resource from poddy server",
			Subcommands: []*cli.Command{
				{
					Name:    "namespace",
					Aliases: []string{"namespaces", "ns", "n"},
					Usage:   "Get namespaces",
					Action:  commands.GetNamespace,
				},
				{
					Name:    "all",
					Aliases: []string{"a"},
					Usage:   "Get all",
					Action:  commands.GetAll,
				},
				{
					Name:    "volume",
					Aliases: []string{"volumes", "vo", "v"},
					Usage:   "Get volumes",
					Action:  commands.GetVolumes,
				},
				{
					Name:    "configmap",
					Aliases: []string{"configmaps", "cm", "c"},
					Usage:   "Get configmaps",
					Action:  commands.GetConfigMap,
				},
			},
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "Delete resource from poddy server",
			Subcommands: []*cli.Command{
				{
					Name:    "configmap",
					Aliases: []string{"configmaps", "cm", "c"},
					Usage:   "Get configmaps",
					Action:  commands.DeleteConfigMap,
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
