package config

import "path/filepath"

// CacheDir is where we should store cached repos
var CacheDir = filepath.Join(ConfigDir, "cache")
