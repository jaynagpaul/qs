package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

// The use command for the CLI
var newCMD = cli.Command{
	Name:        "new",
	Aliases:     []string{"n"},
	Usage:       "Create a default quickstart.toml.",
	Description: "Creates a default quickstart.toml to create a template.",
	UsageText: fmt.Sprintf("%s",
		color.CyanString("qs new"),
	),

	Action: func(ctx *cli.Context) error {
		if _, err := os.Stat("./quickstart.toml"); err == nil {
			return errors.New("quickstart.toml already exists")
		}
		return ioutil.WriteFile("./quickstart.toml", qsTemplate, 0600)
	},
}

var qsTemplate = []byte(`# quickstart.toml example
#
# Refer to https://github.com/jaynagpaul/qs/wiki/quickstart.toml
# for detailed Gopkg.toml documentation.

imports = ["jaynagpaul/license"] # Imported QS files

[[template]]
folder			= "./quickstart" # This cannot be used as a template item
Name			= "NameOfTemplate"	# This cannot be used as a template item
PossibleText 	= ["option1", "option2", "option3"]
Text = "Default"
`)
