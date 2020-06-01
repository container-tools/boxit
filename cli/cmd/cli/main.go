package main

import (
	"fmt"
	"github.com/nicolaferraro/boxnet/cli/cmd/commands"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	var cmd = cobra.Command{
		Use:   "boxnet",
		Short: "Boxnet is a client tool for accessing the boxnet",
	}

	cmd.AddCommand(commands.NewCmdCreate())

	if err := cmd.Execute(); err != nil {
		fmt.Printf("Error occurred: %v\n", err)
		os.Exit(1)
	}
}
