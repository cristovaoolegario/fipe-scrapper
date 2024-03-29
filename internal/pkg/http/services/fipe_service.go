package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/dto"
)

const BASE_URL = "https://veiculos.fipe.org.br"

type Vehicle int

const (
	Car  Vehicle = 1
	Bike Vehicle = 2
)

type FipeService struct {
	Client   *http.Client
	Base_Url string
}

func ProvideDefaultService() *FipeService{
	return NewFipeService(http.Client{
		Timeout: time.Duration(10) * time.Second,
	},BASE_URL)
}

func NewFipeService(client http.Client, base_url string) *FipeService {
	return &FipeService{
		Client:   &client,
		Base_Url: base_url,
	}
}

func setupDataUrl(vehicleType Vehicle, reference, brandCode, modelCode, modelYear, fuelType string) url.Values {
	data := url.Values{}
	data.Set("codigoTipoVeiculo", fmt.Sprint(vehicleType))
	data.Set("codigoTabelaReferencia", reference)
	data.Set("codigoMarca", brandCode)
	data.Set("codigoModelo", modelCode)
	data.Set("anoModelo", modelYear)
	data.Set("codigoTipoCombustivel", fuelType)
	data.Set("tipoConsulta", "tradicional")

	return data
}

func handleJsonBindingError(bodyBytes []byte) error {
	errorObject := dto.FipeError{}
	errBytes := json.Unmarshal(bodyBytes, &errorObject)
	if errBytes != nil {
		return errBytes
	} else {
		return errors.New(errorObject.Erro)
	}
}

func (f *FipeService) setupRequest(resource string, data url.Values) ([]byte, error) {
	encodedData := data.Encode()
	req, err := http.NewRequest("POST", fmt.Sprint(f.Base_Url, "/api/veiculos", resource), strings.NewReader(encodedData))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Referer", BASE_URL)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := f.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("problem with request")
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

func (f *FipeService) GetAllReferences() ([]dto.Referencia, error) {
	responseObject := []dto.Referencia{}
	bodyBytes, _ := f.setupRequest("/ConsultarTabelaDeReferencia", url.Values{})
	err := json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		bindingError := handleJsonBindingError(bodyBytes)
		return nil, bindingError
	}
	return responseObject, nil
}

func (f *FipeService) GetLatestReference() (string, error) {
	resp, err := f.GetAllReferences()
	if err != nil {
		return "", err
	}
	return fmt.Sprint(resp[0].Codigo), nil
}

func (f *FipeService) GetBrands(vehicleType Vehicle) ([]dto.Marca, error) {
	ltsRef, _ := f.GetLatestReference()
	responseObject := []dto.Marca{}
	data := setupDataUrl(vehicleType, ltsRef, "", "", "", "")

	bodyBytes, err := f.setupRequest("/ConsultarMarcas", data)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		bindingError := handleJsonBindingError(bodyBytes)
		return nil, bindingError
	}
	return responseObject, nil
}

func (f *FipeService) GetBrandModels(vehicleType Vehicle, brandCode string) (*dto.MarcaModelo, error) {
	ltsRef, _ := f.GetLatestReference()
	responseObject := &dto.MarcaModelo{}
	data := setupDataUrl(vehicleType, ltsRef, brandCode, "", "", "")

	bodyBytes, err := f.setupRequest("/ConsultarModelos", data)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err == nil && len(responseObject.Modelos) == 0 && len(responseObject.Anos) == 0 {
		bindingError := handleJsonBindingError(bodyBytes)
		return nil, bindingError
	}
	return responseObject, nil
}

func (f *FipeService) GetBrandModelsYears(vehicleType Vehicle, brandCode, modelCode string) ([]dto.Ano, error) {
	ltsRef, _ := f.GetLatestReference()
	responseObject := []dto.Ano{}

	data := setupDataUrl(vehicleType, ltsRef, brandCode, modelCode, "", "")

	bodyBytes, err := f.setupRequest("/ConsultarAnoModelo", data)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		bindingError := handleJsonBindingError(bodyBytes)
		return nil, bindingError
	}

	return responseObject, nil
}

func (f *FipeService) GetVehicle(vehicleType Vehicle, brandCode, modelCode, modelYear, fuelType string) (*dto.Vehicle, error) {
	ltsRef, _ := f.GetLatestReference()
	responseObject := &dto.Vehicle{}

	data := setupDataUrl(vehicleType, ltsRef, brandCode, modelCode, modelYear, fuelType)

	bodyBytes, err := f.setupRequest("/ConsultarValorComTodosParametros", data)

	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err == nil && responseObject.CodigoFipe == "" {
		bindingError := handleJsonBindingError(bodyBytes)
		return nil, bindingError
	}

	return responseObject, nil
}
