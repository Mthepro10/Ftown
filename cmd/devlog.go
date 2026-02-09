/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"time"

	"github.com/spf13/cobra"
)

type Post struct {
	ID              int       `json:"id"`
	Body            string    `json:"body"`
	CommentsCount   int       `json:"comments_count"`
	DurationSeconds int       `json:"duration_seconds"`
	LikesCount      int       `json:"likes_count"`
	ScrapbookURL    string    `json:"scrapbook_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Comments        []struct {
		ID     int `json:"id"`
		Author struct {
			ID          int    `json:"id"`
			DisplayName string `json:"display_name"`
			Avatar      string `json:"avatar"`
		} `json:"author"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"comments"`
}

// devlogCmd represents the devlog command
var devlogCmd = &cobra.Command{
	Use:   "devlog",
	Short: "the root command for devlogs",
	Long:  `the root command for devlogs`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(devlogCmd)
}
