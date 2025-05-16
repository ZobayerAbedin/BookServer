package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yourapp",
	Short: "A simple CLI using Cobra",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from root command!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
