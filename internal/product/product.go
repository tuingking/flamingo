package product

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        int64          `json:"id" orm:"autoIncrement;primaryKey"`
	Barcode   sql.NullString `json:"barcode"`
	Name      string         `json:"name"`
	Slug      string         `json:"slug"`
	ImageURL  string         `json:"image_url" orm:"index"`
	BuyPrice  float64        `json:"buy_price"`
	SellPrice float64        `json:"sell_price"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// GetProductParam adalah parameter untuk get
type GetProductParam struct {
	Name sql.NullString `param:"name" db:"name"`
	Slug sql.NullString `param:"slug" db:"slug"`

	Page    int64    `json:"page"`
	Limit   int64    `json:"limit"`
	ShortBy []string `json:"short_by"`
}
