/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

const (
	credentialService = "ftown"
	credentialUser    = "api_key"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth [API_KEY]",
	Short: "API authentication",
	Long:  `API authentication; find your key in Flavortown settings (the last option)`,
	Args:  cobra.ExactArgs(1), // ensures one argument is provided
	RunE: func(cmd *cobra.Command, args []string) error {
		apiKey := args[0]

		// Save API key in Windows Credential Manager
		if err := keyring.Set(credentialService, credentialUser, apiKey); err != nil {
			return fmt.Errorf("failed to store API key: %w", err)
		}

		client := &http.Client{}
		if verify {
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
		}

		fmt.Println("API key stored securely in Windows Credential Manager ✅")
		return nil
	},
}

var verify bool

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().BoolVarP(&verify, "verify", "v", false, "verify the token")
}
