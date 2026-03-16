package models

//   id integer PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
//   iso_code TEXT NOT NULL UNIQUE,
//   country_name TEXT NOT NULL

type CountryModel struct {
	ID          int    `json:"id"`
	IsoCode     string `json:"iso_code"`
	CountryName string `json:"country_name"`
}
