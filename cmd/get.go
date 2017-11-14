package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jaynagpaul/qs/pkg/qs"
	"github.com/urfave/cli"
)

// The use command for the CLI
var getCMD = cli.Command{
	Name:        "get",
	Aliases:     []string{"g"},
	Usage:       "Install the template but do not run.",
	Description: "This will only install the template but won't run it.",
	UsageText: fmt.Sprintf("Examples:\n\t\tGithub: %s\n\t\tGit Repo: %s\n\t\tStandard Library: %s",
		color.CyanString("qs get jaynagpaul/qs-license"),
		color.CyanString("qs get https://github.com/jaynagpaul/qs-license"),
		color.CyanString("qs get license"),
	),

	Action: func(ctx *cli.Context) error {
		if !ctx.Args().Present() {
			return cli.NewExitError("No path passed", 1)
		}

		path := ctx.Args().First()

		_, err := qs.Get(path)

		return err
	},
}
