package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type EntityRequest struct {
	ID          string    `db:"id"`
	DisplayID   int       `db:"display_id"`
	CompanyID   string    `db:"company_id"`
	StoreID     string    `db:"store_id"`
	WarehouseID *string   `db:"warehouse_id"`
	StatusID    int       `db:"status_id"`
	RequestKind int       `db:"request_kind"`
	CreatedBy   string    `db:"created_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type EntityRequestItem struct {
	ID                string  `db:"id"`
	RequestID         string  `db:"request_id"`
	RequestedQuantity float32 `db:"requested_quantity"`
	MaxQuantity       float32 `db:"max_quantity"`
}

type EntityRequestHistory struct {
	ID          string                  `db:"id"`
	RequestID   string                  `db:"request_id"`
	StatusID    int                     `db:"status_id"`
	ChangedBy   RequestHistoryChangedBy `db:"changed_by"`
	Observation string                  `db:"observation"`
	CreatedAt   time.Time               `db:"created_at"`
}

// RequestHistoryChangedBy is a custom type for JSONB changed_by field
type RequestHistoryChangedBy struct {
	Name             string `json:"name"`
	OrganizationName string `json:"organizationName"`
}

// Value implements the driver.Valuer interface for database storage
func (r RequestHistoryChangedBy) Value() (driver.Value, error) {
	return json.Marshal(r)
}

// Scan implements the sql.Scanner interface for database retrieval
func (r *RequestHistoryChangedBy) Scan(src interface{}) error {
	var data []byte
	switch v := src.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	default:
		return fmt.Errorf("cannot scan %T into RequestHistoryChangedBy", src)
	}
	return json.Unmarshal(data, r)
}
