/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"encoding/json"
	"fmt"
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

		req1, err := http.NewRequest("GET", "https://flavortown.hackclub.com/api/v1/users/me", nil)
		if err != nil {
			return err
		}
		req1.Header.Set("Authorization", "Bearer "+apiKey)
		req1.Header.Set("X-Flavortown-Ext-10376", "true")

		respMe, err := client.Do(req1)
		if err != nil {
			return err
		}
		defer respMe.Body.Close()

		if respMe.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", respMe.Status)
		}

		var me Response
		dec := json.NewDecoder(respMe.Body)
		if err := dec.Decode(&me); err != nil {
			return err
		}

		reqUser, err := http.NewRequest("GET", fmt.Sprintf("https://flavortown.hackclub.com/api/v1/users/%d", me.ID), nil)
		if err != nil {
			return err
		}
		reqUser.Header.Set("Authorization", "Bearer "+apiKey)
		reqUser.Header.Set("X-Flavortown-Ext-10376", "true")

		respUser, err := client.Do(reqUser)
		if err != nil {
			return err
		}
		defer respUser.Body.Close()

		if respUser.StatusCode != http.StatusOK {
			return fmt.Errorf("API error: %s", respUser.Status)
		}

		var user UserResponse
		dec = json.NewDecoder(respUser.Body)
		if err := dec.Decode(&user); err != nil {
			return err
		}

		table, err := LoadTable()
		if err != nil {
			return err
		}

		cookies := 0
		if user.Cookies != nil {
			cookies = *user.Cookies
		}

		switch table {
		case "old":
			fmt.Println("Account info")
			fmt.Println("------------")
			fmt.Println("ID:", user.ID)
			fmt.Println("Display name:", user.DisplayName)
			fmt.Println("Slack ID:", user.SlackID)
			fmt.Println("Projects:", user.ProjectIDs)
			fmt.Println("Cookies:", cookies)
			fmt.Println("Votes:", user.VoteCount)
			fmt.Println("Likes:", user.LikeCount)
			fmt.Printf("Devlog hours: %.1f\n", float64(user.DevlogSecondsTotal)/3600)
			fmt.Printf("Devlog hours today: %.1f\n", float64(user.DevlogSecondsToday)/3600)

		case "modern":
			fmt.Println("Account info")
			fmt.Println("------------")
			fmt.Printf("%s (ID: %d) — Slack: %s\n", user.DisplayName, user.ID, user.SlackID)
			fmt.Printf("Projects: %v | Cookies: %d | Votes: %d | Likes: %d\n", user.ProjectIDs, cookies, user.VoteCount, user.LikeCount)
			fmt.Printf("Devlog: total %.1fh | today %.1fh\n",
				float64(user.DevlogSecondsTotal)/3600, float64(user.DevlogSecondsToday)/3600)

		case "future":
			fmt.Println(">>> ACCOUNT INFO <<<")
			fmt.Printf("[USER-%04d] %s\n", user.ID, user.DisplayName)
			fmt.Printf("SLACK: %s\n", user.SlackID)
			fmt.Printf("PROJECTS (%d): %v\n", len(user.ProjectIDs), user.ProjectIDs)
			fmt.Printf("COOKIES:%4d  VOTES:%4d  LIKES:%4d\n", cookies, user.VoteCount, user.LikeCount)
			fmt.Printf("DEVLOG: TOTAL=%.2fh | TODAY=%.2fh\n",
				float64(user.DevlogSecondsTotal)/3600, float64(user.DevlogSecondsToday)/3600)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(accountInfoCmd)
}
