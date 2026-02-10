/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func saveState(path string, s State) error {
	data, err := yaml.Marshal(s)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600)
}

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config [path]",
	Short: "Set config file for tables",
	Long:  `Set, show, or delete the config file path used by ftown`,
	RunE: func(cmd *cobra.Command, args []string) error {
		stateFile, err := getStateFile()
		if err != nil {
			return err
		}

		if deleteConfig {
			if err := os.Remove(stateFile); err != nil && !os.IsNotExist(err) {
				return err
			}
			fmt.Println("Config path deleted")
			return nil
		}

		if showConfig {
			state, err := loadState(stateFile)
			if err != nil {
				return errors.New("no config path set")
			}
			fmt.Println(state.ConfigPath)
			return nil
		}

		if len(args) == 0 {
			return errors.New("config path required")
		}

		path := args[0]
		if _, err := os.Stat(path); err != nil {
			return fmt.Errorf("config file does not exist: %s", path)
		}

		state := State{
			ConfigPath: path,
		}

		if err := saveState(stateFile, state); err != nil {
			return err
		}

		fmt.Println("Config path saved:", path)
		return nil
	},
}
var deleteConfig bool
var showConfig bool

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVar(&deleteConfig, "delete", false, "Delete saved config path")
	configCmd.Flags().BoolVar(&showConfig, "show", false, "Show saved config path")
}
