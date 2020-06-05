package main

import (
	"fmt"
	"github.com/container-tools/boxit/cli/cmd/commands"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	var cmd = cobra.Command{
		Use:   "boxit",
		Short: "Boxit is a client tool for accessing the boxit network",
	}

	cmd.AddCommand(commands.NewCmdCreate())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}
