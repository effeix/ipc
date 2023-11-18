package ipc

import (
	"os"

	"github.com/effeix/ipc/pkg/ipc"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
    Use:   "explain",
    Short: "Short description.",
    Args:  cobra.ExactArgs(1),
    Run:   func(cmd *cobra.Command, args []string) {
        ipc.ExplainNetwork(args[0])
    },
    PersistentPreRun: func(cmd *cobra.Command, args []string) {
        if !ipc.IsValidCIDR(args[0]) {
            os.Exit(1)
        }
    },
}

func init() {
    rootCmd.AddCommand(explainCmd)
}
