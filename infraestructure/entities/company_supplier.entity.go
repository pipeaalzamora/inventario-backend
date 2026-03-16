package entities

// EntityCompanySupplierWithFiscalData representa el resultado del JOIN entre
// supplier_per_company, supplier y fiscal_data. Incluye toda la información
// relevante del proveedor y sus datos fiscales para mostrar en el frontend.
type EntityCompanySupplierWithFiscalData struct {
	// Campos de supplier
	SupplierID   string `db:"supplier_id"`
	SupplierName string `db:"supplier_name"`
	Description  string `db:"description"`
	Available    bool   `db:"available"`
	CountryID    int    `db:"country_id"`

	// Campos de fiscal_data
	IDFiscal      string `db:"id_fiscal"`
	RawFiscalID   string `db:"raw_fiscal_id"`
	FiscalName    string `db:"fiscal_name"`
	FiscalAddress string `db:"fiscal_address"`
	Email         string `db:"email"`
}
