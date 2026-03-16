package dto

import "sofia-backend/domain/models"

// DtoStoreProduct representa la respuesta de un producto de tienda.
type DtoStoreProduct struct {
	ID            string                          `json:"id"`
	StoreID       string                          `json:"storeId"`
	ProductID     string                          `json:"productId"`
	TagID         int                             `json:"tagId"`
	ProductName   string                          `json:"productName"`
	Image         *string                         `json:"image"`
	ItemSale      bool                            `json:"itemSale"`
	UseRecipe     bool                            `json:"useRecipe"`
	UnitInventory DtoStoreProductUnit             `json:"unitInventory"`
	UnitMatrix    []DtoStoreProductUnitConversion `json:"unitMatrix"`
	Description   string                          `json:"description"`
	Quantities    DtoStoreProductQuantities       `json:"quantities"`
	Suppliers     []DtoStoreProductSupplier       `json:"suppliers,omitempty"`
	CreatedAt     string                          `json:"createdAt"`
	UpdatedAt     string                          `json:"updatedAt"`
}

// DtoStoreProductUnit representa una unidad de medida.
type DtoStoreProductUnit struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Abbreviation string `json:"abbreviation"`
}

// DtoStoreProductUnitConversion representa una conversión de unidad en la matriz.
type DtoStoreProductUnitConversion struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Abbreviation string  `json:"abbreviation"`
	Factor       float32 `json:"factor"`
}

// DtoStoreProductQuantities representa las cantidades del producto de tienda.
type DtoStoreProductQuantities struct {
	MinimalStock float32  `json:"minimalStock"`
	MaximalStock float32  `json:"maximalStock"`
	MaxQuantity  *float32 `json:"maxQuantity"`
}

// DtoStoreProductSupplier representa un proveedor asignado al producto de tienda.
type DtoStoreProductSupplier struct {
	ID          string `json:"id"`
	Name        string `json:"supplierName"`
	Priority    int    `json:"priority"`
	RawFiscalId string `json:"rawFiscalId"`
}

// NewDtoStoreProduct crea un DTO a partir del modelo.
func NewDtoStoreProduct(model *models.ModelProductPerStore) *DtoStoreProduct {
	unitMatrix := make([]DtoStoreProductUnitConversion, len(model.UnitMatrix))
	for i, u := range model.UnitMatrix {
		unitMatrix[i] = DtoStoreProductUnitConversion{
			ID:           u.ID,
			Name:         u.Name,
			Abbreviation: u.Abbreviation,
			Factor:       u.Factor,
		}
	}

	return &DtoStoreProduct{
		ID:          model.ID,
		StoreID:     model.StoreID,
		ProductID:   model.ProductTemplate.ID,
		TagID:       model.TagID,
		ProductName: model.ProductName,
		Image:       model.Image,
		ItemSale:    model.ItemSale,
		UseRecipe:   model.UseRecipe,
		UnitInventory: DtoStoreProductUnit{
			ID:           model.UnitInventory.ID,
			Name:         model.UnitInventory.Name,
			Abbreviation: model.UnitInventory.Abbreviation,
		},
		UnitMatrix:  unitMatrix,
		Description: model.Description,
		Quantities: DtoStoreProductQuantities{
			MinimalStock: model.Quantities.MinimalStock,
			MaximalStock: model.Quantities.MaximalStock,
			MaxQuantity:  model.Quantities.MaxQuantity,
		},
		Suppliers: nil,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

// NewDtoStoreProductSupplier crea un DTO de proveedor a partir del modelo.
// Deprecated: Use enrichSuppliersWithFiscalData en la facade en su lugar.
func NewDtoStoreProductSupplier(model *models.ModelSupplierStoreProduct, supplierName string, productName string) *DtoStoreProductSupplier {
	return &DtoStoreProductSupplier{
		ID:          model.SupplierID,
		Name:        supplierName,
		Priority:    model.Priority,
		RawFiscalId: "",
	}
}
