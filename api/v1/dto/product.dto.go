package dto

type DtoSimpleProduct struct {
	ID            string               `json:"id"`
	Name          string               `json:"name"`
	SKU           string               `json:"sku"`
	CostEstimated float32              `json:"costEstimated"`
	PreviousPrice *float32             `json:"previousPrice,omitempty"`
	Description   *string              `json:"description"`
	Image         *string              `json:"image"`
	Categories    []DtoProductCategory `json:"categories"`
	Codes         []DtoProductCode     `json:"codes"`
}

type DtoProduct struct {
	ID          string               `json:"id"`
	Name        string               `json:"name"`
	Description *string              `json:"description"`
	Image       *string              `json:"image"`
	Categories  []DtoProductCategory `json:"categories"`
	Codes       []DtoProductCode     `json:"codes"`
	UpdatedAt   string               `json:"updatedAt"`
	CreatedAt   string               `json:"createdAt"`
}

type DtoProductCategory struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

type DtoProductCode struct {
	ID    int                `json:"id"`
	Kind  DtoProductCodeKind `json:"kind"`
	Value string             `json:"value"`
}

type DtoProductCodeKind struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
