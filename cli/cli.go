package sha2Cli

import (
	"strings"

	"github.com/petershen0307/getSHA2/core"
	"github.com/urfave/cli"
)

// CreateCli to get cli object
func CreateCli() *cli.App {
	app := cli.NewApp()
	app.Name = "SHA2 generator"
	app.Usage = "Calculate sha2 value"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "path,p",
			Value: "",
			Usage: "input filter path, use ',' for separator",
		},
		cli.StringFlag{
			Name:  "ext,e",
			Value: "",
			Usage: "input filter extension, use ',' for separator",
		},
		cli.StringFlag{
			Name:  "start,s",
			Value: "",
			Usage: "input start path",
		},
	}
	app.Action = cliRouting
	return app
}

func cliRouting(c *cli.Context) error {
	paths := c.String("path")
	exts := c.String("ext")
	startPath := c.String("start")
	core.Start(startPath, strings.Split(paths, ","), strings.Split(exts, ","))
	return nil
}
