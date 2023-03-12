package dto

type Referencia struct {
	Codigo int    `json:"Codigo,omitempty"`
	Mes    string `json:"Mes,omitempty"`
}

type Marca struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}
