/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"github.com/spf13/cobra"
)

type Project struct {
	ID            int    `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	RepoURL       string `json:"repo_url"`
	DemoURL       string `json:"demo_url"`
	ReadmeURL     string `json:"readme_url"`
	AIDeclaration string `json:"ai_declaration"`
	ShipStatus    string `json:"ship_status"`
	DevlogIDs     []int  `json:"devlog_ids"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
}

// projectCmd represents the project command
var projectCmd = &cobra.Command{
	Use:   "project",
	Short: "root command for projects",
	Long:  `root command for ptojects`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectCmd)
}
