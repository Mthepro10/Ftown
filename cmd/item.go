/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

// itemCmd represents the item command
var itemCmd = &cobra.Command{
	Use:   "item [name] [where]",
	Short: "search item by name",
	Long: `search shop item by name
			!!!in this versions typos are prohibeted, but you can write however you want with capital or small letters: hackducky = HaCkducky = HACKDUCKY = HackDucky(original name)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		name_up := args[0]
		name := strings.ToLower(strings.TrimSpace(name_up))
		where := args[1]

		apiKey, err := keyring.Get(credentialService, credentialUser)
		if err != nil {
			return fmt.Errorf("not authenticated, run `ftown auth <API_KEY>` first")
		}

		client := &http.Client{}

		reqMe, err := http.NewRequest("GET", "https://flavortown.hackclub.com/api/v1/store", nil)
		if err != nil {
			return err
		}
		reqMe.Header.Set("Authorization", "Bearer "+apiKey)
		reqMe.Header.Set("X-Flavortown-Ext-10376", "true")

		respMe, err := client.Do(reqMe)
		if err != nil {
			return err
		}
		defer respMe.Body.Close()

		if respMe.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", respMe.Status)
		}

		var resp []Shop
		dec := json.NewDecoder(respMe.Body)
		if err := dec.Decode(&resp); err != nil {
			return err
		}

		var found *Shop
		for i := range resp {
			if strings.ToLower(strings.TrimSpace(resp[i].Name)) == name {
				found = &resp[i]
				break
			}
		}

		if found == nil {
			return fmt.Errorf("item not found")
		}

		url := fmt.Sprintf("https://flavortown.hackclub.com/api/v1/store/%d", found.ID)

		reqItem, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		reqItem.Header.Set("Authorization", "Bearer "+apiKey)
		reqItem.Header.Set("X-Flavortown-Ext-10376", "true")

		itemResp, err := client.Do(reqItem)
		if err != nil {
			return err
		}
		defer itemResp.Body.Close()

		if itemResp.StatusCode != http.StatusOK {
			return fmt.Errorf("%s", itemResp.Status)
		}

		var ans Shop
		dec = json.NewDecoder(itemResp.Body)
		if err := dec.Decode(&ans); err != nil {
			return err
		}

		fmt.Println("id:", ans.ID)
		fmt.Println("name:", ans.Name)
		fmt.Println("description:", ans.LongDescription)
		fmt.Println("stock:", ans.Stock)
		fmt.Println("max quantity:", ans.MaxQty)
		fmt.Println("limited:", ans.Limited)

		switch where {
		case "au":
			fmt.Println("Cost:", ans.TicketCost.AU, "enabled:", ans.Enabled.EnabledAU)
		case "ca":
			fmt.Println("Cost:", ans.TicketCost.CA, "enabled:", ans.Enabled.EnabledCA)
		case "eu":
			fmt.Println("Cost:", ans.TicketCost.EU, "enabled:", ans.Enabled.EnabledEU)
		case "in":
			fmt.Println("Cost:", ans.TicketCost.IN, "enabled:", ans.Enabled.EnabledIN)
		case "uk":
			fmt.Println("Cost:", ans.TicketCost.UK, "enabled:", ans.Enabled.EnabledUK)
		case "us":
			fmt.Println("Cost:", ans.TicketCost.US, "enabled:", ans.Enabled.EnabledUS)
		case "xx":
			fmt.Println("Cost:", ans.TicketCost.XX, "enabled:", ans.Enabled.EnabledXX)
		}

		return nil
	},
}

func init() {
	shopCmd.AddCommand(itemCmd)
}
