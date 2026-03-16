package entities

import "time"

type EntitySupplierOC struct {
	ID         string     `json:"id" db:"id"`
	PurchaseID string     `json:"purchase_id" db:"purchase_id"`
	TokenHash  string     `json:"token_hash" db:"token_hash"`
	Exp        *time.Time `json:"exp" db:"exp_time"`
	Used       bool       `json:"used" db:"used"`
}
