package domain

type Currency struct {
	Code        string  `json:"currency_code"`
	Title       string  `json:"title"`
	Unit        uint    `json:"unit"`
	BuyingRate  float64 `json:"buying_rate"`
	SellingRate float64 `json:"selling_rate"`
}
