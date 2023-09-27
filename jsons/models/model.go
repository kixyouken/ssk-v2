package models

type ModelJson struct {
	Name    string    `json:"name"`
	Table   string    `json:"table"`
	Columns []Columns `json:"columns"`
}

type Columns struct {
	Field string `json:"field"`
}
