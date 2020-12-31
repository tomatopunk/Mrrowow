package main

import (
	"catTable/src/pkg/cmd/cmdRoot"
)

func main() {
	cmdRoot := cmdRoot.CreateCmdRoot()
	cmdRoot.Execute()
}
