package git

import (
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/jaynagpaul/qs/pkg/config"
)

func gitRun(path string, args ...string) (string, error) {
	cmd := exec.Command("git", args...)

	if path != "" {
		cmd.Dir = path
	}
	b, err := cmd.CombinedOutput()
	return string(b), err
}

// Clone the repo to the pkg location.
// Will automatically format username/repo to git repo URL
// If repo is already there, it will only pull.
// Returns the pkgDir and any error that occured
func Clone(path string) (string, error) {
	match, err := regexp.MatchString("[^\n]+/[^\n]+", path)
	var pkgDir string

	// Not a github path
	if err != nil || !match {

		u, err := url.Parse(path)

		if err != nil {
			return "", err
		}
		fmt.Println(u)

		if u.Host == "github.com" {
			// This is done so the cache is reused if they get
			// jaynagpaul/qs-license
			// and https://github.com/jaynagpaul/qs-license
			p := u.Path
			p = p[1:] // Remove first / in path
			pkgDir = p
		} else {
			path = u.String()
			pkgDir = path
		}
	} else {
		pkgDir = path
		path = "https://github.com/" + path
	}

	pkgDir = filepath.Join(config.PkgDir, pkgDir)

	// Repo exists only pull
	if _, err := os.Stat(pkgDir); err == nil {
		// Overwrite changes
		o, err := gitRun(pkgDir, "fetch", "--all")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		o, err = gitRun(pkgDir, "reset", "--hard", "origin/master")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		o, err = gitRun(pkgDir, "pull", "origin", "master")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		return pkgDir, nil
	}

	// clone to repo in cache
	o, err := gitRun(pkgDir, "clone", path)

	if err != nil {
		return "", fmt.Errorf("While cloning: %s\n%s", o, err)
	}

	return pkgDir, nil
}
