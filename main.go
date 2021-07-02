package main

import (
	"easy-search/cmd"
	"easy-search/pkgs/path"
)

func main() {
	path.SetupPath()
	cmd.Execute()
}
