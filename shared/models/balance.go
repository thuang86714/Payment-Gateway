package models

//type balance works as a schema for db Auto-migration
type Balance struct {
    Currency string  `gorm:"primaryKey" json:"currency"`
    Amount   float64 `json:"amount"`
}
