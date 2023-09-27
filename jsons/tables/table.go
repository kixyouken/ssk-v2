package tables

type TableJson struct {
	Name     string   `json:"name"`
	Model    string   `json:"model"`
	Orders   []Orders `json:"orders"`
	Page     int      `json:"page"`
	Limit    int      `json:"limit"`
	Paginate string   `json:"paginate"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}
