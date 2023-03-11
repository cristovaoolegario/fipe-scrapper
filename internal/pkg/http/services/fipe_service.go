package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/dto"
)

const BASE_URL = "https://veiculos.fipe.org.br"

type FipeService struct {
	Client *http.Client
}

func setupRequest(resource string) (*http.Request, error) {
	req, err := http.NewRequest("POST", fmt.Sprint(BASE_URL, "/api/veiculos", resource), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Referer", BASE_URL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (f *FipeService) GetAllReferences() ([]dto.Referencia, error) {
	responseObject := []dto.Referencia{}
	req, err := setupRequest("/ConsultarTabelaDeReferencia")
	if err != nil {
		return nil, err
	}
	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("problem with the request")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(bodyBytes, &responseObject)
	return responseObject, nil
}
