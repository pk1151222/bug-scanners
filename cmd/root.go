package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "bugscanner",
    Short: "A tool for scanning web vulnerabilities",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Bug Scanner is running")
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
