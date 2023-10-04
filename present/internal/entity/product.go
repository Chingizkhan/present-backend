package entity

import "github.com/gofrs/uuid"

type Product struct {
	Id       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Brand    string    `json:"brand"`
	Category string    `json:"category"`
	Currency string    `json:"currency"`
	Quantity int       `json:"quantity"`
	Price    int       `json:"price"`
	OldPrice int       `json:"old_price"`
}
