package models

import "time"

type ModelProduct struct {
	ID            string                 `json:"id" bson:"id"`
	Name          string                 `json:"name" bson:"name"`
	SKU           string                 `json:"sku" bson:"sku"`
	CostEstimated float32                `json:"costEstimated" bson:"cost_estimated"`
	PreviousPrice *float32               `json:"previousPrice,omitempty" bson:"previous_price"`
	Description   *string                `json:"description" bson:"description"`
	Image         *string                `json:"image" bson:"image"`
	Categories    []ModelProductCategory `json:"categories" bson:"categories"`
	Codes         []ModelProductCode     `json:"codes" bson:"codes"`
	UpdatedAt     time.Time              `json:"updatedAt" bson:"updated_at"`
	CreatedAt     time.Time              `json:"createdAt" bson:"created_at"`
}

type ModelProductCategory struct {
	ID          int64  `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Available   bool   `json:"available" bson:"available"`
}

type ModelProductCode struct {
	ID    int                  `json:"id" bson:"id"`
	Kind  ModelProductCodeKind `json:"kind" bson:"kind"`
	Value string               `json:"value" bson:"value"`
}

type ModelProductCodeKind struct {
	ID          int    `json:"id" bson:"id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
}

type ModelProductWithSuppliers struct {
	ModelProduct
	Suppliers []ModelSupplier `json:"suppliers" bson:"suppliers"`
}
