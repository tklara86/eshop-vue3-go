package product

import "github.com/Rhymond/go-money"

type Product struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Image            string `json:"image"`
	Description      string `json:"description"`
	PriceVATExcluded Amount `json:"price_vat_excluded"`
	VAT              Amount `json:"vat"`
	TotalPrice       Amount `json:"total_price"`
}

type Amount struct {
	Money   *money.Money `json:"money"`
	Display string       `json:"display"`
}
