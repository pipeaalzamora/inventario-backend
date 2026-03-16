package entities

import "time"

type EntityProduct struct {
	ID            string    `db:"id"`
	Name          string    `db:"product_name"`
	SKU           string    `db:"sku"`
	Description   *string   `db:"description"`
	Image         *string   `db:"image"`
	CostEstimated float32   `db:"cost_estimated"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

type EntityProductCategory struct {
	ID          int64  `db:"id"`
	Name        string `db:"category_name"`
	Description string `db:"description"`
	Available   bool   `db:"available"`
}

type EntityProductPerCategory struct {
	ID         int64  `db:"id"`
	ProductId  string `db:"product_id"`
	CategoryId int    `db:"category_id"`
}

type EntityCodeKind struct {
	ID          int    `db:"id"`
	Name        string `db:"code_name"`
	Description string `db:"description"`
}

type EntityProductCode struct {
	ID        string `db:"id"`
	ProductId string `db:"product_id"`
	KindId    int    `db:"kind_id"`
	Value     string `db:"code_value"`
}

// EntityProductRequestRestriction moved to product_per_store.entity.go

type EntityProductCodeWithKind struct {
	ID          int    `db:"kind_id"`
	Name        string `db:"code_name"`
	Description string `db:"description"`
	Value       string `db:"code_value"`
}
