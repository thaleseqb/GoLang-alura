package models

import "github.com/google/uuid"

type Pizza struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Price float64   `json:"price"`
}
