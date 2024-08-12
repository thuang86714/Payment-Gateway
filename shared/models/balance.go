package models

type Balance struct {
    Currency string  `gorm:"primaryKey" json:"currency"`
    Amount   float64 `json:"amount"`
}