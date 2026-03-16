package entities

type EntityFiscalData struct {
	ID            string `db:"id"`
	IDFiscal      string `db:"id_fiscal"`
	RawFiscalID   string `db:"raw_fiscal_id"`
	FiscalName    string `db:"fiscal_name"`
	FiscalAddress string `db:"fiscal_address"`
	FiscalState   string `db:"fiscal_state"`
	FiscalCity    string `db:"fiscal_city"`
	Email         string `db:"email"`
}
