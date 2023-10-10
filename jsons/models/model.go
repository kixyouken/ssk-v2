package models

type ModelJson struct {
	Name        string        `json:"name"`
	Table       string        `json:"table"`
	Joins       []Joins       `json:"joins"`
	JoinsCount  []JoinsCount  `json:"joins_count"`
	JoinsSum    []JoinsSum    `json:"joins_sum"`
	Withs       []Withs       `json:"withs"`
	WithsGroups []WithsGroups `json:"withs_groups"`
	Columns     []Columns     `json:"columns"`
	Wheres      []Wheres      `json:"wheres"`
	WheresOr    [][]WheresOr  `json:"wheres_or"`
	Orders      []Orders      `json:"orders"`
	Deleteds    []Deleteds    `json:"deleteds"`
	Groups      []Groups      `json:"groups"`
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

type JoinsSum struct {
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

type WithsGroups struct {
	Table   string    `json:"table"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Type    string    `json:"type"`
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

type WheresOr struct {
	Field string `json:"field"`
	Match string `json:"match"`
	Value string `json:"value"`
}

type Deleteds struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type Groups struct {
	Columns []Columns `json:"columns"`
	Type    string    `json:"type"`
	Group   Group     `json:"group"`
}

type Group struct {
	Field string `json:"field"`
}
