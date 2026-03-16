package entities

import "time"

// CREATE TABLE store (
//
//	id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
//	company_id UUID NOT NULL,
//	store_name character varying NOT NULL,
//	store_address character varying NOT NULL,
//	description character varying NOT NULL,
//	id_cost_center character varying,
//	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
//	updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
//	FOREIGN KEY (company_id) REFERENCES company(id) ON DELETE RESTRICT -- Block deletion
//
// );
type EntityStore struct {
	ID           string    `db:"id"`
	CompanyID    string    `db:"company_id"`
	StoreName    string    `db:"store_name"`
	StoreAddress string    `db:"store_address"`
	Description  string    `db:"description"`
	IDCostCenter string    `db:"id_cost_center"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
}

/*
type EntitySupplierApplied struct {
	SupplierID   string `db:"supplier_id"`
	SupplierName string `db:"supplier_name"`
	Available    bool   `db:"available"`
}
*/
