/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type State struct {
	ConfigPath string `yaml:"config_path"`
}

func getStateFile() (string, error) {
	dir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	appDir := filepath.Join(dir, "ftown")
	if err := os.MkdirAll(appDir, 0700); err != nil {
		return "", err
	}

	return filepath.Join(appDir, "state.yaml"), nil
}

func loadState(path string) (State, error) {
	var s State
	data, err := os.ReadFile(path)
	if err != nil {
		return s, err
	}
	err = yaml.Unmarshal(data, &s)
	return s, err
}

type UserConfig struct {
	Table string `yaml:"table"`
}

func LoadTable() (string, error) {
	stateFile, err := getStateFile()
	if err != nil {
		return "", err
	}

	state, err := loadState(stateFile)
	if err != nil || state.ConfigPath == "" {
		return "old", nil
	}

	data, err := os.ReadFile(state.ConfigPath)
	if err != nil {
		return "old", nil
	}

	var cfg UserConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return "old", nil
	}

	switch cfg.Table {
	case "modern", "old", "future":
		return cfg.Table, nil
	default:
		return "old", nil
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ftown",
	Short: "Extension for Flavortown",
	Long: `


							
   			███████╗████████╗ ██████╗ ██╗    ██╗███╗   ██╗
   			██╔════╝╚══██╔══╝██╔═══██╗██║    ██║████╗  ██║
   			█████╗     ██║   ██║   ██║██║ █╗ ██║██╔██╗ ██║
   			██╔══╝     ██║   ██║   ██║██║███╗██║██║╚██╗██║
   			██║        ██║   ╚██████╔╝╚███╔███╔╝██║ ╚████║
   			╚═╝        ╚═╝    ╚═════╝  ╚══╝╚══╝ ╚═╝  ╚═══╝


		This includes:
			-API authentification: ftown auth [API_KEY]
                --To verify your key you should include the --verify, -v flag
				--You can also logout using the: ftown logout
			-You can see account info using: ftown account-info
			-You can see the shop: ftown shop [where]
			-You can see specific item from shop: ftown shop item [name] [where], if name if longer than 1 word put it in "..."
			-Added project commands
				--You can update projects by name or by id: ftown project update [ID]
				--You can create projects: ftown project create --title --description
				--You can see project info by name or by id: ftown project info [ID]
				--Flag -n,--byname is not recommended because it is not precise
			-Added devlog command
				--You can now see devlog information using: ftown devlog info [ID]
			-You can config the print using the command: ftown config [path to .yaml file]
				--There are 3 ways pf configuration: default(old), modern, future
		
		More informations on https://ftown.gitbook.io/ftown-docs-1/`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
