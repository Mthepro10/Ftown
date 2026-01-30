/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// logoutCmd represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "API delete",
	Long:  `API delte`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := keyring.Delete(
			credentialService,
			credentialUser,
		)
		if err != nil {
			return err
		}

		fmt.Println("API key removed from this computer")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(logoutCmd)
}
