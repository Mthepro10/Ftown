/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ftown",
	Short: "Extension for Flavortown",
	Long: `This includes:
			-API authentification: ftown auth [API_KEY]
                --To verify your key you should include the --verify, -v flag
				--You can also logout using the: ftown logout
			-You can see account info using: ftown accountInfo
			-You can see the shop: ftown shop [where]
			-You can see specific item from shop: ftown shop item [name] [where], if name if longer than 1 word put it in "..."
			-Added project commands
				--You can update projects by name or by id: ftown project update [ID]
				--You can create projects: ftown project create --title --description
				--You can see project info by name or by id: ftown project info [ID]
				--Flag byname is not recommended because it is not precise
			-Added devlog command
				--You can now see devlog information using: ftown devlog info [ID]
		
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
