package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "advent",
	Short: "advent of code CLI",
	Long:  `using advent of code as an excuse to both strengthen my golang muscles & get familiar with cobra`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings
}
