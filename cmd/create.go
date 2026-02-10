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
		reqMe.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		reqMe.Header.Set("Accept", "application/json")
		reqMe.Header.Set("X-Flavortown-Ext-10376", "true")

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

		table, err := LoadTable()
		if err != nil {
			return err
		}

		switch table {
		case "old":
			fmt.Println("Project info")
			fmt.Println("------------")
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

		case "modern":
			fmt.Println("Project info")
			fmt.Println("------------")
			fmt.Printf("%s (ID: %d)\n", resp.Title, resp.ID)
			fmt.Printf("Description: %s\n", resp.Description)
			fmt.Printf("Repo: %s | Demo: %s | Readme: %s\n", resp.RepoURL, resp.DemoURL, resp.ReadmeURL)
			fmt.Printf("AI Declaration: %s | Ship Status: %s\n", resp.AIDeclaration, resp.ShipStatus)
			fmt.Printf("Devlogs: %v\n", resp.DevlogIDs)
			fmt.Printf("Created: %s | Updated: %s\n", resp.CreatedAt, resp.UpdatedAt)

		case "future":
			fmt.Println(">>> PROJECT INFO <<<")
			fmt.Printf("[%04d] %s\n", resp.ID, resp.Title)
			fmt.Printf("DESCRIPTION: %s\n", resp.Description)
			fmt.Printf("REPO: %s  DEMO: %s  README: %s\n", resp.RepoURL, resp.DemoURL, resp.ReadmeURL)
			fmt.Printf("AI DECL: %s | SHIP STATUS: %s\n", resp.AIDeclaration, resp.ShipStatus)
			fmt.Printf("DEVLOG IDs: %v\n", resp.DevlogIDs)
			fmt.Printf("CREATED: %s | UPDATED: %s\n", resp.CreatedAt, resp.UpdatedAt)
		}

		return nil
	},
}

func init() {
	projectCmd.AddCommand(createCmd)
}
