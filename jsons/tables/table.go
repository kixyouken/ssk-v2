package tables

type TableJson struct {
	Name     string   `json:"name"`
	Model    string   `json:"model"`
	Orders   []Orders `json:"orders"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	Paginate string   `json:"paginate"`
	Wheres   []Wheres `json:"wheres"`
	Withs    []Withs  `json:"withs"`
	Joins    []Joins  `json:"joins"`
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

type Joins struct {
	Model   string    `json:"model"`
	Foreign string    `json:"foreign"`
	Key     string    `json:"key"`
	Join    string    `json:"join"`
	Columns []Columns `json:"columns"`
}
