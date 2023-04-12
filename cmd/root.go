/*
Copyright © 2023 Oni Harnantyo oni.harnantyo97@gmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my-go-clean-architecture",
	Short: "My golang clean architecture",
	Long:  `My golang clean architecture project structure inspired by Uncle Bob.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&configFile, "config", "./config/config.yaml", "Config file")
}
