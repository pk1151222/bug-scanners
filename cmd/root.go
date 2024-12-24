package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bug-scanner",
	Short: "A tool to scan for subdomains and security bugs",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Bug Scanner Tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
