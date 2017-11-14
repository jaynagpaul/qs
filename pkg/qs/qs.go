package qs

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fatih/color"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/jaynagpaul/qs/pkg/git"
	"github.com/pelletier/go-toml"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Get the repo but do not execute it.
func Get(path string) (string, error) {
	w := wow.New(os.Stdout, spin.Get(spin.Dots), " Downloading "+path)
	w.Start()

	pkgPath, err := git.Clone(path)

	if err != nil {
		w.PersistWith(spin.Spinner{Frames: []string{""}}, "")
		return "", err
	}

	w.PersistWith(spin.Spinner{Frames: []string{"👍"}}, " Downloaded "+path+"!")
	return pkgPath, nil
}

// Run will parse the directory at path in search of a quickstart.toml.
// After finding the file it will call ExecFile on it.
func Run(path string) error {
	p, err := findQS(path)

	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(p)

	if err != nil {
		return err
	}

	c := Config{}

	if err := toml.Unmarshal(b, &c); err != nil {
		return err
	}

	if len(c.Templates) == 0 {
		color.Magenta("No template to run!")
		return nil
	}
	tmpls := make([]string, 0, len(c.Templates))
	for _, t := range c.Templates {
		tmpls = append(tmpls, t.Name())
	}

	var tmpl string
	// Select Template
	survey.AskOne(&survey.Select{
		Message: "Choose a template",
		Options: tmpls,
		Default: tmpls[0],
	}, &tmpl, survey.Required)

	// Run all imports
	// TODO, maybe check for cyclic imports
	for _, imp := range c.Imports {
		if p, err := Get(imp); err == nil {
			if err := Run(p); err != nil {
				return err
			}
		} else {
			return err
		}
	}

	if err != nil {
		return err
	}

	// Run at end
	// TODO execute templates
	for _, cmd := range c.Commands {
		fmt.Printf("Running %s", color.HiBlueString(cmd))
		o, err := execCommand(cmd)
		if err != nil {
			return fmt.Errorf("%s: %s\n%s", cmd, o, err)
		}
	}

	return nil

}

func execCommand(cmd string) (string, error) {
	args := strings.Split(cmd, " ")

	if len(args) == 0 {
		return "", nil
	}

	b, err := exec.Command(args[0], args[1:]...).CombinedOutput()

	return string(b), err
}

// Search the directory for a quickstart.toml file
func findQS(path string) (string, error) {
	if p := filepath.Join(path, "quickstart.toml"); fileExists(p) {
		return p, nil
	}

	return "", errors.New("No quickstart.toml in " + path)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

// Config is the structural representation of the quickstart.toml
type Config struct {
	Name      string     `toml:"name"`
	Imports   []string   `toml:"imports"`
	Commands  []string   `toml:"commands"`
	Templates []Template `toml:"template"`
}

// A Template to execute
type Template map[string]interface{}

// Name returns the name of the template, "" if not set
func (t Template) Name() string {
	if name, ok := t["TemplateName"].(string); ok {
		return name
	}
	return ""
}

// Folder returns the folder of the template, "" if not set
func (t Template) Folder() string {
	if name, ok := t["TemplateFolder"].(string); ok {
		return name
	}
	return ""
}
