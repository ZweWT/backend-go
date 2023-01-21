package product

import "time"

// Product represents a single product.
type Product struct {
	ID          string    `db:"product_id" json:"id"`
	Name        string    `db:"name" json:"name"`
	CategoryID  int       `db:"category_id" json:"category_id"`
	Cost        int       `db:"cost" json:"cost"`
	DateCreated time.Time `db:"date_created" json:"date_created"`
	DateUpdated time.Time `db:"date_updated" json:"date_updated"`
}

// NewProduct contains information needed to create a new Product.
type NewProduct struct {
	Name       string `json:"name" validate:"required"`
	CategoryID int    `json:"category_id" validate:"required"`
	Cost       int    `json:"cost" validate:"required"`
}

// UpdateProduct contains information needed to update a Product.
type UpdateProduct struct {
	Name       *string `json:"name"`
	CategoryID *int    `json:"category_id"`
	Cost       *int    `json:"cost"`
}
