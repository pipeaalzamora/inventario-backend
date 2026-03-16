package models

type ModelFiscalData struct {
	ID            string `json:"id"`
	IDFiscal      string `json:"idFiscal"`
	RawFiscalID   string `json:"rawFiscalId"`
	FiscalName    string `json:"fiscalName"`
	FiscalAddress string `json:"fiscalAddress"`
	FiscalState   string `json:"fiscalState"`
	FiscalCity    string `json:"fiscalCity"`
	Email         string `json:"email"`
}

type EconomicActivityClassModel struct {
	ID         int
	CountryID  int
	SystemName string
	SystemCode string
}

type EconomicActivityModel struct {
	ID                      int
	EconomicActivityClassID int
	ActivityName            string
	ActivityCode            string
	Description             string
}
