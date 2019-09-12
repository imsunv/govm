package main

import "github.com/urfave/cli"

type Options struct {
	Env       string
	Classpath string
}

func build(options *Options) *cli.App {

	cliApp := cli.NewApp()
	cliApp.Name = "govm"
	cliApp.Version = "v1"
	cliApp.HideHelp = true

	cliApp.Action = func(c *cli.Context) {
		if c.GlobalBool("help") {
			cli.ShowAppHelp(c)
			cli.OsExiter(1)
		}
	}

	cliApp.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "help, h",
			Usage: "show help",
		},
		cli.StringFlag{
			Name:        "env, e",
			Usage:       "the supported environment",
			Value:       "stable_masterjd",
			Destination: &options.Env,
		},
		cli.StringFlag{
			Name:        "classpath, cp",
			Usage:       "目录和 zip/jar 文件的类搜索路",
			Value:       "",
			Destination: &options.Classpath,
		},
	}

	return cliApp
}
