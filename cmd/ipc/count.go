package ipc

import (
	"fmt"

	"github.com/effeix/ipc/pkg/ipc"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
    Use:   "count",
    Short:  "Counts how many IP addresses fit in the CIDR range.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        res := ipc.CountIPs(args[0])
        fmt.Println(res)
    },
}

func init() {
    rootCmd.AddCommand(countCmd)
}