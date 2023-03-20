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

type Vehicle struct {
	Valor            string `json:"Valor"`
	Marca            string `json:"Marca"`
	Modelo           string `json:"Modelo"`
	AnoModelo        int    `json:"AnoModelo"`
	Combustivel      string `json:"Combustivel"`
	CodigoFipe       string `json:"CodigoFipe"`
	MesReferencia    string `json:"MesReferencia"`
	Autenticacao     string `json:"Autenticacao"`
	TipoVeiculo      int    `json:"TipoVeiculo"`
	SiglaCombustivel string `json:"SiglaCombustivel"`
	DataConsulta     string `json:"DataConsulta"`
}
