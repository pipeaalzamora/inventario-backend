package entities

import (
	"time"
)

type EntityWarehouse struct {
	ID string `db:"id"`
	//CompanyId            string             `db:"company_id"`
	StoreId             string  `db:"store_id"`
	Name                string  `db:"warehouse_name"`
	Description         *string `db:"description"`
	Address             *string `db:"warehouse_address"`
	Phone               *string `db:"warehouse_phone"`
	IsMomeventWarehouse bool    `db:"is_momevent_warehouse"`
	//DeliveryInstructions string             `db:"delivery_instructions"`
	//WorkingHours         EntityWorkingHours `db:"working_hours"`
	//WorkingTimeZone      string             `db:"working_timezone"`
	CreatedAt time.Time `db:"created_at"`
}

/*
	type EntityWorkingHours struct {
		Monday    []string `json:"monday"`
		Tuesday   []string `json:"tuesday"`
		Wednesday []string `json:"wednesday"`
		Thursday  []string `json:"thursday"`
		Friday    []string `json:"friday"`
		Saturday  []string `json:"saturday"`
		Sunday    []string `json:"sunday"`
	}
*/
// EntityWarehousePerProduct representa el stock de un producto por tienda en una bodega.
// Ahora referencia store_product_id (producto por tienda) en lugar de product_company_id.
type EntityWarehousePerProduct struct {
	ID                   string  `json:"id" db:"id"`
	StoreProductId       string  `json:"store_product_id" db:"store_product_id"` // Referencia a product_per_store
	StoreId              string  `json:"store_id" db:"store_id"`                 // Para consultas por tienda
	WarehouseId          string  `json:"warehouse_id" db:"warehouse_id"`
	WarehouseIdReference *string `json:"warehouse_id_reference" db:"warehouse_id_reference"` // Bodega de referencia
	Direction            *string `json:"direction" db:"direction"`                           // Dirección: "IN" o "OUT"
	InStock              float32 `json:"in_stock" db:"in_stock"`
	CostAvg              float32 `json:"cost_avg" db:"cost_avg"` // Costo promedio
	InTransit            float32 `json:"in_transit" db:"in_transit"`
	Ordered              float32 `json:"ordered" db:"ordered"`
}

// Esto es necesario para que el tipo EntityWorkingHours pueda ser utilizado en consultas SQL
/*
func (e *EntityWorkingHours) Scan(value interface{}) error {
	if value == nil {
		*e = EntityWorkingHours{}
		return nil
	}

	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("EntityWorkingHours: expected []byte, got %T", value)
	}

	return json.Unmarshal(b, e)
}

// para inserts/updates
func (e EntityWorkingHours) Value() (driver.Value, error) {
	return json.Marshal(e)
}
*/
