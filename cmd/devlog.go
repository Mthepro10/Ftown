/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// devlogCmd represents the devlog command
var devlogCmd = &cobra.Command{
	Use:   "devlog",
	Short: "the root command for devlogs",
	Long:  `the root command for devlogs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Errorf("coming soon!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devlogCmd)
}
