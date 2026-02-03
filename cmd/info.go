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
	"strconv"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// infoCmd represents the info command
var infoCmd = &cobra.Command{
	Use:   "info [ID]",
	Short: "see info about project",
	Long:  `see info about project by id`,
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
			req.Header.Set("X-Flavortown-Ext-10376", "true")

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

		req, err := http.NewRequest(
			"GET",
			"https://flavortown.hackclub.com/api/v1/projects/"+id,
			nil,
		)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+apiKey)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("X-Flavortown-Ext-10376", "true")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			b, _ := io.ReadAll(resp.Body)
			return fmt.Errorf("fetch failed (%d): %s", resp.StatusCode, string(b))
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
	projectCmd.AddCommand(infoCmd)
}
