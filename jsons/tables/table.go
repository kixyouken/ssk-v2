package tables

type TableJson struct {
	Name     string   `json:"name"`
	Model    string   `json:"model"`
	Orders   []Orders `json:"orders"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	Paginate string   `json:"paginate"`
	Wheres   []Wheres `json:"wheres"`
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
