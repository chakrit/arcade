package main

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "arcade",
	Short: "Starts the arcade simulation.",
}

func init() {
	RootCmd.AddCommand(PingCmd)
}
