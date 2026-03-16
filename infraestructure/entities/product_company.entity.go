package entities

// DEPRECATED: EntityProductCompany ha sido reemplazado por EntityProductPerStore
// Esta entidad se mantiene comentada para referencia durante la migración.
// Ver: product_per_store.entity.go
/*
import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type EntityProductCompany struct {
	ID              string           `db:"id"`
	CompanyID       string           `db:"company_id"`
	ProductID       string           `db:"product_id"`
	TagID           int              `db:"tag_id"`
	SKU             string           `db:"sku"`
	ProductName     string           `db:"product_name"`
	ItemPurchase    bool             `db:"item_purchase"`
	ItemSale        bool             `db:"item_sale"`
	ItemInventory   bool             `db:"item_inventory"`
	IsFrozen        bool             `db:"is_frozen"`
	UseRecipe       bool             `db:"use_recipe"`
	UnitInventoryID int              `db:"unit_inventory_id"`
	UnitInventory   string           `db:"unit_inventory"`
	UnitPurchaseID  int              `db:"unit_purchase_id"`
	UnitPurchase    string           `db:"unit_purchase"`
	UnitMatrix      EntityUnitMatrix `db:"unit_matrix"`
	CostLast        float64          `db:"cost_last"`
	Description     string           `db:"description"`
	CostEstimated   float64          `db:"cost_estimated"`
	CostAvg         float64          `db:"cost_avg"`
	MinimalStock    float64          `db:"minimal_stock"`
	MaximalStock    float64          `db:"maximal_stock"`
	MinimalOrder    float64          `db:"minimal_order"`
	CreatedAt       string           `db:"created_at"`
	UpdatedAt       string           `db:"updated_at"`
}
*/

// EntityUnitMatrix y EntityConversionUnit se usan también en ProductPerStore

/*
type EntityUnitMatrix map[int]EntityConversionUnit

func (e *EntityUnitMatrix) Scan(value any) error {
	if value == nil {
		*e = nil
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("EntityUnitMatrix: unexpected type %T in Scan", value)
	}

	if len(b) == 0 {
		*e = EntityUnitMatrix{}
		return nil
	}

	// Opción 1: si en el JSON las keys son numéricas (no strings)
	var tmp map[int]EntityConversionUnit
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("EntityUnitMatrix: error unmarshaling JSON: %w", err)
	}

	*e = tmp

	return nil
}

func (e EntityUnitMatrix) Value() (driver.Value, error) {
	if e == nil {
		return nil, nil
	}

	b, err := json.Marshal(e)
	if err != nil {
		return nil, fmt.Errorf("EntityUnitMatrix: error marshaling JSON: %w", err)
	}
	return b, nil
}
*/
/*
type EntityUnitOfMeasurement struct {
	ID           string `db:"id"`
	Name         string `db:"name"`
	Abbreviation string `db:"abbreviation"`
	Description  string `db:"description"`
	IsBase       bool   `db:"is_base"`
}

type EntityConversionUnit struct {
	Id           int     `json:"id" jsonb:"id"`
	Name         string  `json:"name" jsonb:"name"`
	Abbreviation string  `json:"abbreviation" jsonb:"abbreviation"`
	Description  string  `json:"description" jsonb:"description"`
	Factor       float32 `json:"factor" jsonb:"factor"`
}
*/
