package models

// BaseProduct represents data about a product without an ID.
type BaseProduct struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Product extends BaseProduct with an ID and InStock fields.
type Product struct {
	ID          string `json:"id"`
	BaseProduct `json:",inline"`
	InStock     bool `json:"inStock"`
}
