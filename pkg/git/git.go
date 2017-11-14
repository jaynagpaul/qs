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

func gitRun(args ...string) (string, error) {
	b, err := exec.Command("git", args...).CombinedOutput()
	return string(b), err
}

// Clone the repo to the cache location.
// Will automatically format username/repo to git repo URL
// If repo is already there, it will only pull.
// Returns the cacheDir and any error that occured
func Clone(path string) (string, error) {
	match, err := regexp.MatchString("[^\n]+/[^\n]+", path)
	var cacheDir string

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
			cacheDir = p
		} else {
			path = u.String()
			cacheDir = path
		}
	} else {
		cacheDir = path
		path = "https://github.com/" + path
	}

	// Repo exists only pull
	if _, err := os.Stat(filepath.Join(config.CacheDir, cacheDir)); err == nil {
		// Overwrite changes
		o, err := gitRun("fetch", "--all")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		o, err = gitRun("reset", "--hard", "origin/master")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		o, err = gitRun("pull", "origin", "master")
		if err != nil {
			return "", fmt.Errorf("While pulling: %s\n%s", o, err)
		}

		return cacheDir, nil
	}

	// clone to repo in cache
	o, err := gitRun("clone", path, filepath.Join(config.CacheDir, cacheDir))

	if err != nil {
		return "", fmt.Errorf("While cloning: %s\n%s", o, err)
	}

	return cacheDir, nil
}
