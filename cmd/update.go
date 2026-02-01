/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [ID]",
	Short: "Update project",
	Long:  `Update project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey, err := keyring.Get(credentialService, credentialUser)
		if err != nil {
			return fmt.Errorf("not authenticated, run `ftown auth <API_KEY>` first")
		}

		client := &http.Client{}
		var id string

		if byname != "" {
			req, err := http.NewRequest(
				"GET",
				"https://flavortown.hackclub.com/api/v1/projects?query="+url.QueryEscape(byname),
				nil,
			)
			if err != nil {
				return err
			}
			req.Header.Set("Authorization", "Bearer "+apiKey)

			resp, err := client.Do(req)
			if err != nil {
				return err
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				b, _ := io.ReadAll(resp.Body)
				return fmt.Errorf("search failed (%d): %s", resp.StatusCode, string(b))
			}

			var search projectSearchResponse
			if err := json.NewDecoder(resp.Body).Decode(&search); err != nil {
				return err
			}

			if len(search.Projects) == 0 {
				return fmt.Errorf("no project found with name %q", byname)
			}

			id = strconv.Itoa(search.Projects[0].ID)
		} else {
			if len(args) == 0 {
				return fmt.Errorf("missing project ID")
			}
			id = args[0]
		}

		data := url.Values{}
		if title != "" {
			data.Set("title", title)
		}
		if description != "" {
			data.Set("description", description)
		}
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

		if len(data) == 0 {
			return fmt.Errorf("no fields provided to update")
		}

		req, err := http.NewRequest(
			"PATCH",
			"https://flavortown.hackclub.com/api/v1/projects/"+id,
			strings.NewReader(data.Encode()),
		)
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("update failed (%d): %s", resp.StatusCode, string(b))
		}

		var project Project
		if err := json.NewDecoder(resp.Body).Decode(&project); err != nil {
			return err
		}

		fmt.Println("Project ID:", project.ID)
		fmt.Println("Title:", project.Title)
		fmt.Println("Description:", project.Description)
		fmt.Println("Repo URL:", project.RepoURL)
		fmt.Println("Demo URL:", project.DemoURL)
		fmt.Println("Readme URL:", project.ReadmeURL)
		fmt.Println("AI Declaration:", project.AIDeclaration)
		fmt.Println("Ship Status:", project.ShipStatus)
		fmt.Println("Devlog IDs:", project.DevlogIDs)
		fmt.Println("Created At:", project.CreatedAt)
		fmt.Println("Updated At:", project.UpdatedAt)

		return nil
	},
}

func init() {
	projectCmd.AddCommand(updateCmd)
}
