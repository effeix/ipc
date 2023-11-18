package ipc

import (
    "github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
    Use:   "convert",
    Short: "Short description.",
    Args:  cobra.ExactArgs(1),
    Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
    rootCmd.AddCommand(convertCmd)
}
