package main

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "arcade",
	Short: "Starts the arcade simulation.",
	RunE: runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	return nil
}


