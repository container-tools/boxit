package main

import (
	"fmt"
	"github.com/container-tools/boxit/server/pkg/commands"
	"os"
)

func main() {
	cmd := commands.NewCmdRoot()

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}
