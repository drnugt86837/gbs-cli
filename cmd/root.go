/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/drnugt86837/gbs-cli/pkg/generator"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gbs-cli",
	Short: "A CLI tool for generating Go project structures",
	Long: `gbs-cli is a command-line interface (CLI) tool written in Go.
It helps you quickly generate the necessary files and folder structures for your Go projects.

It uses the Cobra library, which is a CLI framework for Go applications, to provide an intuitive command-line experience.

Usage:
  gbs-cli [command]

Available Commands:
  create      Create a new module structure with predefined files and directories

Flags:
  -h, --help   help for gbs-cli

Use "gbs-cli [command] --help" for more information about a command.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var createCmd = &cobra.Command{
	Use:   "create [moduleName]",
	Short: "Create folder structure and files",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		moduleName := args[0]

		err := generator.CreateModuleStructure(moduleName)
		if err != nil {
			fmt.Println("Error creating folder structure:", err)
			return
		}
		fmt.Println("Folder structure and files created successfully.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gbs-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(createCmd)
}
