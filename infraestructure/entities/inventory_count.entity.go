package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type EntityInventoryCount struct {
	ID              string                     `db:"id"`
	DisplayID       string                     `db:"display_id"`
	StoreID         string                     `db:"store_id"`
	StoreName       string                     `db:"store_name"`
	CompanyID       string                     `db:"company_id"`
	WarehouseID     string                     `db:"warehouse_id"`
	WarehouseName   string                     `db:"warehouse_name"`
	CreatedBy       string                     `db:"created_by"`
	CreatedByName   string                     `db:"created_by_name"`
	AssignedTo      *string                    `db:"assigned_to"`
	AssignedToName  *string                    `db:"assigned_to_name"`
	Status          string                     `db:"status"`
	ScheduledAt     time.Time                  `db:"scheduled_at"`
	CompletedAt     *time.Time                 `db:"completed_at"`
	CreatedAt       time.Time                  `db:"created_at"`
	UpdatedAt       time.Time                  `db:"updated_at"`
	MovementTrackId string                     `db:"movement_track_id"`
	MetaData        InventoryMetadata          `db:"metadata"`
	CountItems      []EntityInventoryCountItem `db:"-"`
}

type InventoryCountStatus string

const (
	InventoryCountStatusPending   InventoryCountStatus = "pending"
	InventoryCountStatusCompleted InventoryCountStatus = "completed"
	InventoryCountStatusCancelled InventoryCountStatus = "cancelled"
)

type EntityInventoryCountItem struct {
	ID                   string    `db:"id"`
	InventoryCountID     string    `db:"inventory_count_id"`
	ProductID            string    `db:"store_product_id"`
	WarehouseID          string    `db:"warehouse_id"`
	ScheduledAt          time.Time `db:"scheduled_at"`
	ProductName          string    `db:"product_name"`
	ProductSKU           string    `db:"product_sku"`
	ProductImage         *string   `db:"product_image"`
	ProductBaseUnitID    int       `db:"product_base_unit_id"`
	ProductBaseUnitAbv   string    `db:"product_base_unit_abv"`
	IncidenceImageURL    *string   `db:"incidence_image_url"`
	IncidenceObservation *string   `db:"incidence_observation"`
}

type EntityInventoryCountMetadata struct {
	Completed bool    `json:"completed"`
	UnitId    int     `json:"unitId"`
	UnitAbv   string  `json:"unitAbv"`
	Count     float32 `json:"count"`
	Factor    float32 `json:"factor"`
}

type InventoryMetadata map[string][]EntityInventoryCountMetadata

func (m *InventoryMetadata) Scan(value any) error {
	if value == nil {
		*m = nil
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("InventoryMetadata: tipo inesperado %T en Scan", value)
	}

	// Deserializar JSON a tu mapa
	var tmp map[string][]EntityInventoryCountMetadata
	if err := json.Unmarshal(b, &tmp); err != nil {
		return fmt.Errorf("InventoryMetadata: error al hacer Unmarshal: %w", err)
	}

	*m = tmp
	return nil
}

func (m InventoryMetadata) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	b, err := json.Marshal(m)
	if err != nil {
		return nil, fmt.Errorf("InventoryMetadata: error al hacer Marshal: %w", err)
	}
	return b, nil
}
