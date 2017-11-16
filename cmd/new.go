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
# Order of Operations
#   1. imports
#   2. ask all options
#   3. commands
#   4. run templates

	
name = "QS - Quickstart Any Application"

[[template]]
TemplateFolder = "./quickstart" # This is not asked to the user. Default: .
TemplateName = "DefaultName"    # This is not asked to the user. Default: top level name

# Imported QS files. Run in order, and after executing templates.
# Template variables may be used.
imports = ["jaynagpaul/license"] # This is not asked to the user.

# Run first. This is not asked to the user.
# You can use template variables here
commands = ["create-react-app"] 

PossibleOption = ["option1", "option2", "option3"]
DefaultText = "Default"
DefaultNum = 1

`)
