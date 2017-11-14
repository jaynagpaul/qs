package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// The use command for the CLI
var listCMD = cli.Command{
	Name:        "list",
	Aliases:     []string{"ls", "l"},
	Usage:       "List all currently downloaded templates.",
	Description: "Checks the cache dir for all downloaded templates.",
	UsageText: fmt.Sprintf("%s",
		color.CyanString("qs list"),
	),

	Action: func(ctx *cli.Context) error {
		fmt.Println("TODO")
		return nil
	},
}
