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

		table, err := LoadTable()
		if err != nil {
			return err
		}

		switch table {
		case "old":
			fmt.Println("Item info")
			fmt.Println("------------")
			fmt.Println("ID:", ans.ID)
			fmt.Println("Name:", ans.Name)
			fmt.Println("Description:", ans.LongDescription)
			fmt.Println("Stock:", ans.Stock)
			fmt.Println("Max Quantity:", ans.MaxQty)
			fmt.Println("Limited:", ans.Limited)

			switch where {
			case "au":
				fmt.Println("Cost:", ans.TicketCost.AU, "Enabled:", ans.Enabled.EnabledAU)
			case "ca":
				fmt.Println("Cost:", ans.TicketCost.CA, "Enabled:", ans.Enabled.EnabledCA)
			case "eu":
				fmt.Println("Cost:", ans.TicketCost.EU, "Enabled:", ans.Enabled.EnabledEU)
			case "in":
				fmt.Println("Cost:", ans.TicketCost.IN, "Enabled:", ans.Enabled.EnabledIN)
			case "uk":
				fmt.Println("Cost:", ans.TicketCost.UK, "Enabled:", ans.Enabled.EnabledUK)
			case "us":
				fmt.Println("Cost:", ans.TicketCost.US, "Enabled:", ans.Enabled.EnabledUS)
			case "xx":
				fmt.Println("Cost:", ans.TicketCost.XX, "Enabled:", ans.Enabled.EnabledXX)
			}

		case "modern":
			fmt.Println("Item info")
			fmt.Println("------------")
			fmt.Printf("%s (ID: %d)\n", ans.Name, ans.ID)
			fmt.Printf("Description: %s\n", ans.LongDescription)
			fmt.Printf("Stock: %d | Max Qty: %d | Limited: %v\n", ans.Stock, ans.MaxQty, ans.Limited)

			switch where {
			case "au":
				fmt.Printf("Cost: %v | Enabled: %v (AU)\n", ans.TicketCost.AU, ans.Enabled.EnabledAU)
			case "ca":
				fmt.Printf("Cost: %v | Enabled: %v (CA)\n", ans.TicketCost.CA, ans.Enabled.EnabledCA)
			case "eu":
				fmt.Printf("Cost: %v | Enabled: %v (EU)\n", ans.TicketCost.EU, ans.Enabled.EnabledEU)
			case "in":
				fmt.Printf("Cost: %v | Enabled: %v (IN)\n", ans.TicketCost.IN, ans.Enabled.EnabledIN)
			case "uk":
				fmt.Printf("Cost: %v | Enabled: %v (UK)\n", ans.TicketCost.UK, ans.Enabled.EnabledUK)
			case "us":
				fmt.Printf("Cost: %v | Enabled: %v (US)\n", ans.TicketCost.US, ans.Enabled.EnabledUS)
			case "xx":
				fmt.Printf("Cost: %v | Enabled: %v (XX)\n", ans.TicketCost.XX, ans.Enabled.EnabledXX)
			}

		case "future":
			fmt.Println(">>> ITEM INFO <<<")
			fmt.Printf("[%04d] %s\n", ans.ID, ans.Name)
			fmt.Printf("DESCRIPTION: %s\n", ans.LongDescription)
			fmt.Printf("STOCK: %d | MAX QTY: %d | LIMITED: %v\n", ans.Stock, ans.MaxQty, ans.Limited)

			switch where {
			case "au":
				fmt.Printf("COST (AU): %v | ENABLED: %v\n", ans.TicketCost.AU, ans.Enabled.EnabledAU)
			case "ca":
				fmt.Printf("COST (CA): %v | ENABLED: %v\n", ans.TicketCost.CA, ans.Enabled.EnabledCA)
			case "eu":
				fmt.Printf("COST (EU): %v | ENABLED: %v\n", ans.TicketCost.EU, ans.Enabled.EnabledEU)
			case "in":
				fmt.Printf("COST (IN): %v | ENABLED: %v\n", ans.TicketCost.IN, ans.Enabled.EnabledIN)
			case "uk":
				fmt.Printf("COST (UK): %v | ENABLED: %v\n", ans.TicketCost.UK, ans.Enabled.EnabledUK)
			case "us":
				fmt.Printf("COST (US): %v | ENABLED: %v\n", ans.TicketCost.US, ans.Enabled.EnabledUS)
			case "xx":
				fmt.Printf("COST (XX): %v | ENABLED: %v\n", ans.TicketCost.XX, ans.Enabled.EnabledXX)
			}
		}

		return nil
	},
}

func init() {
	shopCmd.AddCommand(itemCmd)
}
