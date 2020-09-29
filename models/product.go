package models

type product struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	UnitMeasurement string `json:"unit"`
	Quantity        int    `json:"quantity"`
	CategoryID      int    `json:"category"`
}
