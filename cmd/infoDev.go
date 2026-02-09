/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// infoDevCmd represents the infoDev command
var infoDevCmd = &cobra.Command{
	Use:   "info [ID]",
	Short: "Used to show devlogs",
	Long:  `Used to show devlogs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		if id == "" {
			return errors.New("ID is required")
		}

		apiKey, err := keyring.Get(credentialService, credentialUser)
		if err != nil {
			return fmt.Errorf("not authenticated, run `ftown auth <API_KEY>` first")
		}

		client := &http.Client{}

		reqMe, err := http.NewRequest("GET", "https://flavortown.hackclub.com/api/v1/devlogs/"+id, nil)
		if err != nil {
			return err
		}
		reqMe.Header.Set("Authorization", "Bearer "+apiKey)
		reqMe.Header.Set("X-Flavortown-Ext-10376", "true")
		reqMe.Header.Set("Accept", "application/json")

		respMe, err := client.Do(reqMe)
		if err != nil {
			return err
		}
		defer respMe.Body.Close()

		if respMe.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", respMe.Status)
		}

		var post Post
		dec := json.NewDecoder(respMe.Body)
		if err := dec.Decode(&post); err != nil {
			return err
		}

		fmt.Println("Devlog:")
		fmt.Println("  ID:", post.ID)
		fmt.Println("  Body: \n ------------------------ \n", post.Body, "\n ------------------------\n")
		fmt.Println("  Comments Count:", post.CommentsCount)
		fmt.Println("  Duration (seconds):", post.DurationSeconds)
		fmt.Println("  Likes Count:", post.LikesCount)
		fmt.Println("  Scrapbook URL:", post.ScrapbookURL)
		fmt.Println("  Created At:", post.CreatedAt)
		fmt.Println("  Updated At:", post.UpdatedAt)
		fmt.Println("  Comments:")
		for _, c := range post.Comments {
			fmt.Println("    Comment ID:", c.ID)
			fmt.Println("      Author ID:", c.Author.ID)
			fmt.Println("      Display Name:", c.Author.DisplayName)
			fmt.Println("      Avatar:", c.Author.Avatar)
			fmt.Println("      Body:", c.Body)
			fmt.Println("      Created At:", c.CreatedAt)
			fmt.Println("      Updated At:", c.UpdatedAt)
		}

		return nil
	},
}

func init() {
	devlogCmd.AddCommand(infoDevCmd)
}
