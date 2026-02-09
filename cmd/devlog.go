/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// devlogCmd represents the devlog command
var devlogCmd = &cobra.Command{
	Use:   "devlog",
	Short: "the root command for devlogs",
	Long:  `the root command for devlogs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devlogCmd)
}
