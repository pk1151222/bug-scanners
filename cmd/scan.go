package cmd

import (
    "github.com/spf13/cobra"
    "github.com/yourusername/bug-scanner/internal/scanner"
)

var scanCmd = &cobra.Command{
    Use:   "scan",
    Short: "Run the bug scanner",
    Run: func(cmd *cobra.Command, args []string) {
        scanner.StartScan(args)
    },
}

func init() {
    rootCmd.AddCommand(scanCmd)
}
