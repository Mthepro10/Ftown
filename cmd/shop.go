/*
Copyright © 2026 MIHAI DRAGHICI <mihaidraghiici023@gmail.com>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

type Shop struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	OldPrices           []any  `json:"old_prices"`
	Limited             bool   `json:"limited"`
	Stock               int    `json:"stock"`
	Type                string `json:"type"`
	ShowInCarousel      bool   `json:"show_in_carousel"`
	AccessoryTag        string `json:"accessory_tag"`
	AghContents         any    `json:"agh_contents"`
	AttachedShopItemIDs []int  `json:"attached_shop_item_ids"`
	BuyableBySelf       bool   `json:"buyable_by_self"`
	LongDescription     string `json:"long_description"`
	MaxQty              int    `json:"max_qty"`
	OnePerPersonEver    bool   `json:"one_per_person_ever"`
	SalePercentage      int    `json:"sale_percentage"`
	ImageURL            string `json:"image_url"`

	Enabled struct {
		EnabledAU bool `json:"enabled_au"`
		EnabledCA bool `json:"enabled_ca"`
		EnabledEU bool `json:"enabled_eu"`
		EnabledIN bool `json:"enabled_in"`
		EnabledUK bool `json:"enabled_uk"`
		EnabledUS bool `json:"enabled_us"`
		EnabledXX bool `json:"enabled_xx"`
	} `json:"enabled"`

	TicketCost struct {
		BaseCost float64 `json:"base_cost"`
		AU       float64 `json:"au"`
		CA       float64 `json:"ca"`
		EU       float64 `json:"eu"`
		IN       float64 `json:"in"`
		UK       float64 `json:"uk"`
		US       float64 `json:"us"`
		XX       float64 `json:"xx"`
	} `json:"ticket_cost"`
}

// shopCmd represents the shop command
var shopCmd = &cobra.Command{
	Use:   "shop [where]",
	Short: "show shop",
	Long:  `show entire shop`,
	RunE: func(cmd *cobra.Command, args []string) error {
		where := args[0]

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

		userBody, err := io.ReadAll(respMe.Body)
		if err != nil {
			return err
		}

		var ans []Shop
		if err := json.Unmarshal(userBody, &ans); err != nil {
			return err
		}

		for _, item := range ans {
			fmt.Println("------------")
			fmt.Println("id:", item.ID)
			fmt.Println("name:", item.Name)
			fmt.Println("description:", item.LongDescription)
			fmt.Println("stock:", item.Stock)
			fmt.Println("max quantity:", item.MaxQty)
			fmt.Println("limited:", item.Limited)

			switch where {
			case "au":
				fmt.Println("Cost:", item.TicketCost.AU, "enabled:", item.Enabled.EnabledAU)
			case "ca":
				fmt.Println("Cost:", item.TicketCost.CA, "enabled:", item.Enabled.EnabledCA)
			case "eu":
				fmt.Println("Cost:", item.TicketCost.EU, "enabled:", item.Enabled.EnabledEU)
			case "in":
				fmt.Println("Cost:", item.TicketCost.IN, "enabled:", item.Enabled.EnabledIN)
			case "uk":
				fmt.Println("Cost:", item.TicketCost.UK, "enabled:", item.Enabled.EnabledUK)
			case "us":
				fmt.Println("Cost:", item.TicketCost.US, "enabled:", item.Enabled.EnabledUS)
			case "xx":
				fmt.Println("Cost:", item.TicketCost.XX, "enabled:", item.Enabled.EnabledXX)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(shopCmd)
}
