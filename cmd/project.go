/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "root command for projects",
	Long:  `root command for ptojects`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
