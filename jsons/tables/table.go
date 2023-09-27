package tables

type TableJson struct {
	Name   string   `json:"name"`
	Model  string   `json:"model"`
	Orders []Orders `json:"orders"`
	Page   string   `json:"page"`
}

type Orders struct {
	Field string `json:"field"`
	Sort  string `json:"sort"`
}
