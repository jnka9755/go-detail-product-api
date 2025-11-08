package models

import "time"

type Product struct {
	ID                string      `json:"id"`
	Title             string      `json:"title"`
	Price             float64     `json:"price"`
	CurrencyID        string      `json:"currency_id"`
	Condition         string      `json:"condition"`
	AvailableQuantity int         `json:"available_quantity"`
	SoldQuantity      int         `json:"sold_quantity"`
	Pictures          []Picture   `json:"pictures"`
	Attributes        []Attribute `json:"attributes"`
	CategoryID        string      `json:"category_id"`
	Category          Category    `json:"category"`
	SellerID          string      `json:"seller_id"`
	Seller            Seller      `json:"seller"`
	Shipping          Shipping    `json:"shipping"`
	CreatedAt         time.Time   `json:"created_at"`
}

type Picture struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Attribute struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Shipping struct {
	Mode         string `json:"mode"`
	FreeShipping bool   `json:"free_shipping"`
	StorePickUp  bool   `json:"store_pick_up"`
}

type Seller struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Reputation float64 `json:"reputation"`
	Location   string  `json:"location"`
}

type Category struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	ParentID *string `json:"parent_id,omitempty"`
}
