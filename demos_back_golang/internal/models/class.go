package models

import (
	"github.com/shopspring/decimal"
)

type Class struct {
	ID        int64           `json:"id"`
	DateTime  string          `json:"datetime"`
	Location  string          `json:"location"`
	Price     decimal.Decimal `json:"price"`
	TrainerID int64           `json:"trainer_id"`
	CreatedAt string          `json:"created_at"`
	UpdatedAT string          `json:"updated_at"`
}
