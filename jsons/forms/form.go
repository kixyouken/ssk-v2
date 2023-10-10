package forms

type FormJson struct {
	Name   string   `json:"name"`
	Model  string   `json:"model"`
	Wheres []Wheres `json:"wheres"`
	Withs  []Withs  `json:"withs"`
}

type Withs struct {
	Model   string    `json:"model"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
	Has     string    `json:"has"`
	Wheres  []Wheres  `json:"wheres"`
}

type Columns struct {
	Field  string `json:"field"`
	Format string `json:"format"`
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
