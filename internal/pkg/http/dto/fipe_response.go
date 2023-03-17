package dto

type FipeError struct {
	Codigo string `json:"codigo,omitempty"`
	Erro   string `json:"erro,omitempty"`
}

type Referencia struct {
	Codigo int    `json:"Codigo,omitempty"`
	Mes    string `json:"Mes,omitempty"`
}

type Marca struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

type Modelo struct {
	Label string `json:"Label,omitempty"`
	Value int    `json:"Value,omitempty"`
}

type Ano struct {
	Label string `json:"Label,omitempty"`
	Value string `json:"Value,omitempty"`
}

type MarcaModelo struct {
	Modelos []Modelo `json:"Modelos,omitempty"`
	Anos    []Ano    `json:"Anos,omitempty"`
}
