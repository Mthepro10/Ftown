/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/

package cmd

import (
	"fmt"

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

		fmt.Println("API key stored securely in Windows Credential Manager ✅")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
