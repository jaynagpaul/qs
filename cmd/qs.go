package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/jaynagpaul/qs/pkg/git"
	"github.com/urfave/cli"
)

// App is the CLI App
var App = cli.NewApp()

func init() {
	App.Name = "qs"
	App.Usage = "Quickstart an application to avoid writing repetitive code."
	App.Version = "1.0.0"
	App.UsageText = "\n\tqs <path>\n\tqs <command> [options]\n\n" + fmt.Sprintf("Examples:\n\t\tGithub: %s\n\t\tGit Repo: %s\n\t\tStandard Library: %s",
		color.CyanString("qs use jaynagpaul/qs-license"),
		color.CyanString("qs use https://github.com/jaynagpaul/qs-license"),
		color.CyanString("qs use license"),
	)

	App.Action = func(ctx *cli.Context) error {
		if !ctx.Args().Present() {
			return cli.NewExitError("No path passed", 1)
		}

		w := wow.New(os.Stdout, spin.Get(spin.Dots), " Downloading")
		w.Start()

		// path to git repo
		// two options:
		// jaynagpaul/qs => https://github.com/jaynagpaul/qs
		// https://github.com/jaynagpaul/qs => https://github.com/jaynagpaul/qs
		path := ctx.Args().First()

		path, err := git.Clone(path)

		if err != nil {
			w.PersistWith(spin.Spinner{Frames: []string{""}}, "")
			return err
		}

		w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Downloaded!")

		//qs.Run()
		return nil
	}

	App.Commands = []cli.Command{
		getCMD,
		listCMD,
		newCMD,
	}

	App.ExitErrHandler = func(ctx *cli.Context, err error) {
		if err == nil {
			return
		}

		fmt.Printf("%s:\n%s", color.HiRedString("Error"), err.Error())
	}

	App.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "no-color",
			Usage:       "No colored output",
			EnvVar:      "QS_COLORBLIND",
			Hidden:      false,
			Destination: &color.NoColor,
		},
	}

	cli.AppHelpTemplate = `{{.Name}}{{if .Usage}} - {{.Usage}}{{end}}

Usage: {{.UsageText}}

Commands:{{range .VisibleCategories}}{{if .Name}}
	{{.Name}}:{{end}}{{range .VisibleCommands}}
	{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}
	 
Options: {{range .VisibleFlags}}
	{{.}}{{end}}
	`

	cli.CommandHelpTemplate = `{{.Name}} - {{.Usage}}

Description:
	{{.Description}}

Usage:
	{{if .UsageText}}{{.UsageText}}{{else}}{{.Name}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}
 
Options:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Aliases}}

Aliases: {{join .Aliases ", "}}{{end}}
	`
}
