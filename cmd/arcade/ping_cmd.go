package main

import (
	"github.com/spf13/cobra"
	"github.com/chakrit/arcade"
	"github.com/chakrit/arcade/engine"
	"context"
	"time"
)

var PingCmd = &cobra.Command{
	Use:   "ping (address) [address...]",
	Short: "test node connectivity",
	RunE:  runPingCmd,
}

func runPingCmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return arcade.ErrPrecondition.
			WithMessage("at least one node address required")
	}

	node, ctx := engine.NewNode(args[0]), context.Background()
	if err := node.Introspect(ctx); err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	if err := node.Ping(ctx, 10); err != nil {
		return err
	}

	return nil
}
