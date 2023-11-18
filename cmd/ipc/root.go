package ipc

import (
	"os"

	"github.com/spf13/cobra"
)

var version string = "0.0.1"

var rootCmd = &cobra.Command{
    Use:  "ipc",
    Short: "ipc - IP address and subnet calculator",
	Version: version,
    Long: ``,
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) == 0 {
            cmd.Help()
            os.Exit(0)
        }
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}