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

type projectSearchResponse struct {
	Projects []Project `json:"projects"`
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
var byname string

func init() {
	rootCmd.AddCommand(projectCmd)
	//
	createCmd.Flags().StringVarP(&title, "title", "t", "", "project title")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "project description")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "", "project repo")
	createCmd.Flags().StringVarP(&demo, "demo", "e", "", "project demo")
	createCmd.Flags().StringVarP(&readme, "readme", "m", "", "project readme")
	createCmd.Flags().StringVarP(&ai_declaration, "ai_declaration", "a", "", "project ai_declaration")
	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")

	//
	infoCmd.Flags().StringVarP(&byname, "byname", "n", "", "command will pass argument id searching the project just by name")

	//
	updateCmd.Flags().StringVarP(&title, "title", "t", "", "project title")
	updateCmd.Flags().StringVarP(&description, "description", "d", "", "project description")
	updateCmd.Flags().StringVarP(&repo, "repo", "r", "", "project repo")
	updateCmd.Flags().StringVarP(&demo, "demo", "e", "", "project demo")
	updateCmd.Flags().StringVarP(&readme, "readme", "m", "", "project readme")
	updateCmd.Flags().StringVarP(&ai_declaration, "ai_declaration", "a", "", "project ai_declaration")
	updateCmd.Flags().StringVarP(&byname, "byname", "n", "", "command will pass argument id updating the project just by name")
}
