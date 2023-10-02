package models

type ModelJson struct {
	Name       string       `json:"name"`
	Table      string       `json:"table"`
	Joins      []Joins      `json:"joins"`
	Withs      []Withs      `json:"withs"`
	WithsCount []WithsCount `json:"Withs_count"`
	Columns    []Columns    `json:"columns"`
	Wheres     []Wheres     `json:"wheres"`
}

type Columns struct {
	Field  string `json:"field"`
	Format string `json:"format"`
}

type Joins struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Columns []Columns `json:"columns"`
	Wheres  []Wheres  `json:"wheres"`
}

type Withs struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
	Has     string    `json:"has"`
	Wheres  []Wheres  `json:"wheres"`
}

type WithsCount struct {
	Table   string   `json:"model"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Wheres  []Wheres `json:"wheres"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}

type Wheres struct {
	Field string `json:"field"`
	Match string `json:"match"`
	Value string `json:"value"`
}
