/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

type Response struct {
	ID int `json:"id"`
}

type UserResponse struct {
	ID                 int    `json:"id"`
	SlackID            string `json:"slack_id"`
	DisplayName        string `json:"display_name"`
	Avatar             string `json:"avatar"`
	ProjectIDs         []int  `json:"project_ids"`
	Cookies            *int   `json:"cookies"`
	VoteCount          int    `json:"vote_count"`
	LikeCount          int    `json:"like_count"`
	DevlogSecondsTotal int    `json:"devlog_seconds_total"`
	DevlogSecondsToday int    `json:"devlog_seconds_today"`
}

var accountInfoCmd = &cobra.Command{
	Use:   "account-info",
	Short: "Show full account info",
	RunE: func(cmd *cobra.Command, args []string) error {

		apiKey, err := keyring.Get(credentialService, credentialUser)
		if err != nil {
			return fmt.Errorf("not authenticated, run `ftown auth <API_KEY>` first")
		}

		client := &http.Client{}

		reqMe, err := http.NewRequest("GET", "https://flavortown.hackclub.com/api/v1/users/me", nil)
		if err != nil {
			return err
		}
		reqMe.Header.Set("Authorization", "Bearer "+apiKey)

		respMe, err := client.Do(reqMe)
		if err != nil {
			return err
		}
		defer respMe.Body.Close()

		if respMe.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", respMe.Status)
		}

		meBody, err := io.ReadAll(respMe.Body)
		if err != nil {
			return err
		}

		var me Response
		if err := json.Unmarshal(meBody, &me); err != nil {
			return err
		}

		reqUser, err := http.NewRequest(
			"GET",
			fmt.Sprintf("https://flavortown.hackclub.com/api/v1/users/%d", me.ID),
			nil,
		)
		if err != nil {
			return err
		}
		reqUser.Header.Set("Authorization", "Bearer "+apiKey)

		respUser, err := client.Do(reqUser)
		if err != nil {
			return err
		}
		defer respUser.Body.Close()

		if respUser.StatusCode != http.StatusOK {
			return fmt.Errorf("API error: %s", respUser.Status)
		}

		userBody, err := io.ReadAll(respUser.Body)
		if err != nil {
			return err
		}

		var user UserResponse
		if err := json.Unmarshal(userBody, &user); err != nil {
			return err
		}

		fmt.Println("Account info")
		fmt.Println("------------")
		fmt.Println("ID:", user.ID)
		fmt.Println("Display name:", user.DisplayName)
		fmt.Println("Slack ID:", user.SlackID)
		fmt.Println("Projects:", user.ProjectIDs)
		fmt.Println("Cookies:", *user.Cookies)
		fmt.Println("Votes:", user.VoteCount)
		fmt.Println("Likes:", user.LikeCount)
		fmt.Printf("Devlog hours: %.1f\n", float64(user.DevlogSecondsTotal)/3600)
		fmt.Printf("Devlog hours today: %.1f\n", float64(user.DevlogSecondsToday)/3600)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountInfoCmd)
}
