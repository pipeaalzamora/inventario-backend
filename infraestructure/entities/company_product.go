package entities

type CompanyProduct struct {
	Id            string    `db:"id"`
	CompanyId     string    `db:"company_id"` //UUID
	ProductId     string    `db:"product_id"` //UUID
	TagId         int       `db:"tag_id"`
	SKU           string    `db:"sku"`
	ProductName   string    `db:"product_name"`
	ItemPurchase  bool      `db:"item_purchase"`
	ItemSale      bool      `db:"item_sale"`
	ItemInventory bool      `db:"item_inventory"`
	IsFrozen      bool      `db:"is_frozen"`
	UseRecipe     bool      `db:"use_recipe"`
	UnitPurchase  UnitValue `db:"unit_purchase"`
	UnitInventory UnitValue `db:"unit_inventory"`
	CostLast      float32   `db:"cost_last"`
	Description   string    `db:"description"`
	CostEstimated float32   `db:"cost_estimated"`
	CostAvg       float32   `db:"cost_avg"`
	MinimalStock  float32   `db:"minimal_stock"`
	MaximalStock  float32   `db:"maximal_stock"`
}

type UnitValue string

const (
	AT      UnitValue = "AT"
	BALDE   UnitValue = "BALDE"
	BANDEJA UnitValue = "BANDEJA"
	BOLSAS  UnitValue = "BOLSAS"
	BOTELLA UnitValue = "BOTELLA"
	FRASCO  UnitValue = "FRASCO"
	KILOS   UnitValue = "KILOS"
	LITROS  UnitValue = "LITROS"
	PAQUETE UnitValue = "PAQUETE"
	PORCION UnitValue = "PORCION"
	SACO    UnitValue = "SACO"
	UNIDAD  UnitValue = "UNIDAD"
)
