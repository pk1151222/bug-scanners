package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"bug-scanners/internal/scanner"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan for subdomains and vulnerabilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting the scan...")
		// Call your scanning logic here
		scanner.Scan("example.com")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}

