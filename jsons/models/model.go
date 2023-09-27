package models

type ModelJson struct {
	Name    string    `json:"name"`
	Table   string    `json:"table"`
	Joins   []Joins   `json:"joins"`
	Columns []Columns `json:"columns"`
}

type Columns struct {
	Field string `json:"field"`
}

type Joins struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Columns []Columns `json:"columns"`
}
