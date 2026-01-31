/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a new project",
	Long:  `create a new project`,
	RunE: func(cmd *cobra.Command, args []string) error {

		apiKey, err := keyring.Get(credentialService, credentialUser)
		if err != nil {
			return fmt.Errorf("not authenticated, run `ftown auth <API_KEY>` first")
		}

		data := url.Values{}
		data.Set("title", title)
		data.Set("description", description)
		if repo != "" {
			data.Set("repo_url", repo)
		}
		if demo != "" {
			data.Set("demo_url", demo)
		}
		if readme != "" {
			data.Set("readme_url", readme)
		}
		if ai_declaration != "" {
			data.Set("ai_declaration", ai_declaration)
		}

		client := &http.Client{}

		reqMe, err := http.NewRequest("POST", "https://flavortown.hackclub.com/api/v1/projects", strings.NewReader(data.Encode()))
		if err != nil {
			return err
		}
		reqMe.Header.Set("Authorization", "Bearer "+apiKey)

		respMe, err := client.Do(reqMe)
		if err != nil {
			return err
		}
		defer respMe.Body.Close()
		if respMe.StatusCode != http.StatusCreated {
			return fmt.Errorf("%s", respMe.Status)
		}

		userBody, err := io.ReadAll(respMe.Body)
		if err != nil {
			return err
		}

		var resp Project
		if err := json.Unmarshal(userBody, &resp); err != nil {
			return err
		}

		fmt.Println("Project ID: ", resp.ID)
		fmt.Println("Title: ", resp.Title)
		fmt.Println("Description: ", resp.Description)
		fmt.Println("Repo URL: ", resp.RepoURL)
		fmt.Println("Demo URL: ", resp.DemoURL)
		fmt.Println("Readme URL: ", resp.ReadmeURL)
		fmt.Println("AI Declaration: ", resp.AIDeclaration)
		fmt.Println("Ship Status: ", resp.ShipStatus)
		fmt.Println("Devlog IDs: ", resp.DevlogIDs)
		fmt.Println("Created At: ", resp.CreatedAt)
		fmt.Println("Updated At: ", resp.UpdatedAt)

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
	projectCmd.AddCommand(createCmd)
	createCmd.Flags().StringVarP(&title, "title", "t", "", "project title")
	createCmd.Flags().StringVarP(&description, "description", "d", "", "project description")
	createCmd.Flags().StringVarP(&repo, "repo", "r", "", "project repo")
	createCmd.Flags().StringVarP(&demo, "demo", "e", "", "project demo")
	createCmd.Flags().StringVarP(&readme, "readme", "m", "", "project readme")
	createCmd.Flags().StringVarP(&ai_declaration, "ai_declaration", "a", "", "project ai_declaration")

	createCmd.MarkFlagRequired("title")
	createCmd.MarkFlagRequired("description")
}
