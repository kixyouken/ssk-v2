package models

type ModelJson struct {
	Name    string    `json:"name"`
	Table   string    `json:"table"`
	Joins   []Joins   `json:"joins"`
	Withs   []Withs   `json:"withs"`
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

type Withs struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
	Has     string    `json:"has"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}
