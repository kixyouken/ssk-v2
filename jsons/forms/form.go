package forms

type FormJson struct {
	Name       string       `json:"name"`
	Model      string       `json:"model"`
	Withs      []Withs      `json:"withs"`
	WithsCount []WithsCount `json:"withs_count"`
}

type Withs struct {
	Model   string    `json:"model"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Columns []Columns `json:"columns"`
	Orders  []Orders  `json:"orders"`
	Has     string    `json:"has"`
}

type Columns struct {
	Field  string `json:"field"`
	Format string `json:"format"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}

type WithsCount struct {
	Model   string   `json:"model"`
	Foreign string   `json:"foreign"`
	Key     string   `json:"key"`
	Wheres  []Wheres `json:"wheres"`
}

type Wheres struct {
	Field string `json:"field"`
	Match string `json:"match"`
	Value string `json:"value"`
}
