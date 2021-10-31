// +build darwin

package config

import (
	"os"
	"path/filepath"
)

var (
	homeDirPath, _        = os.UserHomeDir()
	DefaultLoggerFilePath = filepath.Join(homeDirPath, ".ztat/logs/core.log")
	DefaultLoggerLevel    = "debug"
)
