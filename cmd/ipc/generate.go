package ipc

import (
	"fmt"
	"os"

	"github.com/effeix/ipc/pkg/ipc"
	"github.com/spf13/cobra"
)

var binary bool
var oneline bool
var outputFile string

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short:  "Generates the full list of IP addresses included in the range of a given CIDR block.",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        numIPs := ipc.CountIPs(args[0])
        if numIPs > 100000 {
            cmd.Printf("This operation will generate %d IP addresses and might take a while. Are you sure you want to continue? (y/N) ", numIPs)
            
            var confirmation string
            fmt.Scanln(&confirmation)
            if confirmation != "y" {
                cmd.Println("Aborting.")
                return
            }
        }

        if outputFile != "" {
            file, err := os.Create(outputFile)
            if err != nil {
                cmd.Println(err)
                return
            }
            defer file.Close()
            cmd.SetOutput(file)
        }

        ips := ipc.GetPrintable(
            ipc.GenerateIPs(args[0]),
            binary,
            oneline,
        )

        for _, ip := range ips {
            cmd.Print(ip)
        }
    },
}

func init() {
    generateCmd.Flags().BoolVarP(&binary, "binary", "b", false, "Output the binary representations of the IPs.")
    generateCmd.Flags().BoolVarP(&oneline, "oneline", "l", false, "Output the list of IPs in a single line, separated by commas.")
    generateCmd.Flags().StringVarP(&outputFile, "output", "o", "", "Output file")
    rootCmd.AddCommand(generateCmd)
}