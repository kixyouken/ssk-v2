package models

type ModelJson struct {
	Name       string       `json:"name"`
	Table      string       `json:"table"`
	Joins      []Joins      `json:"joins"`
	JoinsCount []JoinsCount `json:"joins_count"`
	Withs      []Withs      `json:"withs"`
	WithsCount []WithsCount `json:"withs_count"`
	WithsSum   []WithsSum   `json:"withs_sum"`
	Columns    []Columns    `json:"columns"`
	Wheres     []Wheres     `json:"wheres"`
	Orders     []Orders     `json:"orders"`
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

type JoinsCount struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Wheres  []Wheres  `json:"wheres"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
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
	Table   string   `json:"table"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Wheres  []Wheres `json:"wheres"`
}

type WithsSum struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Wheres  []Wheres  `json:"wheres"`
	Columns []Columns `json:"columns"`
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
