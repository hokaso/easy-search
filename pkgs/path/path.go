package path

import (
	"os"
	"path"
	"runtime"
)

var ConfigPath string

// Tests will load config and fixture from project root path.
func init() {
	_, filename, _, _ := runtime.Caller(0)
	ConfigPath = path.Join(path.Dir(filename), "../..")
}

// This is used for command line to use running path as root for config path.
func SetupPath() {
	ConfigPath, _ = os.Getwd()
}
