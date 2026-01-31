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

var title string
var description string
var repo string
var demo string
var readme string
var ai_declaration string

func init() {
	rootCmd.AddCommand(projectCmd)
	createCmd.Flags().StringVarP(&title, "title", "t", "", "project title")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "project description")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "", "project repo")
	createCmd.Flags().StringVarP(&demo, "demo", "e", "", "project demo")
	createCmd.Flags().StringVarP(&readme, "readme", "m", "", "project readme")
	createCmd.Flags().StringVarP(&ai_declaration, "ai_declaration", "a", "", "project ai_declaration")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")
}
