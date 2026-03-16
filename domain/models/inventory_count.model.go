package models

import "time"

// Cambiar modelo (eliminar storeName, wharehouseName, createdByName, assignedToName y agregarlos al dto)
type ModelInventoryCount struct {
	ID              string                        `json:"id"`
	DisplayID       string                        `json:"displayId"`
	StoreID         string                        `json:"storeId"`
	StoreName       string                        `json:"storeName"`
	CompanyID       string                        `json:"companyId"`
	WarehouseID     string                        `json:"warehouseId"`
	WarehouseName   string                        `json:"warehouseName"`
	CreatedBy       string                        `json:"createdBy"`
	CreatedByName   string                        `json:"createdByName"`
	AssignedTo      *string                       `json:"assignedTo"`
	AssignedToName  *string                       `json:"assignedToName"`
	Status          string                        `json:"status"`
	ScheduledAt     time.Time                     `json:"scheduledAt"`
	CompletedAt     *time.Time                    `json:"completedAt"`
	CreatedAt       time.Time                     `json:"createdAt"`
	UpdatedAt       time.Time                     `json:"updatedAt"`
	MovementTrackId string                        `json:"movementTrackId"`
	CountItems      []ModelInventoryCountItem     `json:"countItems"`
	Metadata        []ModelInventoryCountMetadata `json:"metaData"`
}

type ModelInventoryCountItem struct {
	ProductID            string    `json:"productId"`
	ProductName          string    `json:"productName"`
	ProductSKU           string    `json:"productSku"`
	ProductImage         *string   `json:"productImage"`
	UnitId               int       `json:"unitId"`
	UnitAbv              string    `json:"unitAbv"`
	ScheduledAt          time.Time `json:"scheduledAt"`
	IncidenceImageURL    *string   `json:"incidenceImageUrl"`
	IncidenceObservation *string   `json:"incidenceObservation"`
}

type ModelInventoryCountMetadata struct {
	ProductID  string                     `json:"productId"`
	Completed  bool                       `json:"completed"`
	Total      float32                    `json:"total"`
	UnitsCount []ModelInventoryUnitsCount `json:"unitsCount"`
}

type ModelInventoryUnitsCount struct {
	UnitId  int     `json:"unitId"`
	UnitAbv string  `json:"unitAbv"`
	Count   float32 `json:"count"`
	Factor  float32 `json:"factor"`
}

///////////////// Extensions /////////////////////

func (m *ModelInventoryCountMetadata) GetUnitsMap() map[int]ModelInventoryUnitsCount {
	unitsMap := make(map[int]ModelInventoryUnitsCount)
	for _, unitCount := range m.UnitsCount {
		unitsMap[unitCount.UnitId] = unitCount
	}
	return unitsMap
}
