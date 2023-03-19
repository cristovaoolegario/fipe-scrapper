package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const mockError = `{"codigo": "2","erro": "Parâmetros inválidos"}`
const mockReferences = `[{"Codigo": 295,"Mes": "março/2023 "},{"Codigo": 294,"Mes": "fevereiro/2023 "}]`
const mockBrands = `[{"Label": "Acura","Value": "1"},{"Label": "Agrale","Value": "2"}]`
const mockBrandsModels = `{"Modelos":[{"Label":"116iA 1.6 TB 16V 136cv 5p","Value":6146},{"Label":"118i M Sport 1.5 TB 12V Aut. 5p","Value":9955}],"Anos":[{"Label":"32000 Gasolina","Value":"32000-1"},{"Label":"2023 Gasolina","Value":"2023-1"}]}`
const mockModelYears = `[{"Label": "2020 Gasolina","Value": "2020-1"},{"Label": "2019 Gasolina","Value": "2019-1"}]`

func mockHandler(mockData string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mockResponse := strings.NewReader(mockData)
		byteValue, _ := ioutil.ReadAll(mockResponse)

		w.Header().Add("Content-Type", "application/json")
		w.Write(byteValue)
	}
}

func TestFipeService_GetAllReferences(t *testing.T) {
	t.Run("Should return References data when response is valid", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(mockHandler(mockReferences)))

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetAllReferences()

		if err != nil && response == nil {
			t.Errorf("Shouldn't have an error, got: %s", err)
		}
	})

	t.Run("Should return error when response is invalid", func(t *testing.T) {
		expected := "Parâmetros inválidos"
		ts := httptest.NewServer(http.HandlerFunc(mockHandler(mockError)))

		defer ts.Close()
		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetAllReferences()

		if err == nil && response != nil {
			t.Errorf("Should have an error, got %d responses", len(response))
		}
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}

func TestFipeService_GetLatestReference(t *testing.T) {
	t.Run("Should return the latest reference code when response is valid", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(mockHandler(mockReferences)))

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetLatestReference()

		if err != nil && response == "" {
			t.Errorf("Shouldn't have an error, got: %s", err)
		}
		if response != "295" {
			t.Errorf("Expected code to be 295, got:%s", response)
		}
	})

	t.Run("Should return error when response is invalid", func(t *testing.T) {
		expected := "Parâmetros inválidos"
		ts := httptest.NewServer(http.HandlerFunc(mockHandler(mockError)))

		defer ts.Close()
		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetLatestReference()

		if err == nil && response != "" {
			t.Errorf("Should have an error, got response: %s", response)
		}
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}

func TestFipeService_GetBrands(t *testing.T) {
	t.Run("Should return all Brands when response is valid", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarMarcas", mockHandler(mockBrands))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrands(Car)

		if err != nil && response == nil {
			t.Errorf("Shouldn't have an error, got: %s", err.Error())
		}
		if len(response) != 2 {
			t.Errorf("Expected response to have 2 items, got: %d ", len(response))
		}
	})

	t.Run("Should return error when response is invalid", func(t *testing.T) {
		expected := "Parâmetros inválidos"
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarMarcas", mockHandler(mockError))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrands(Car)

		if err == nil && response != nil {
			t.Errorf("Should have an error, got response: %s", response)
		}
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}

func TestFipeService_GetBrandModels(t *testing.T) {
	t.Run("Should return All Models from Brands when the brand id is valid", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarModelos", mockHandler(mockBrandsModels))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrandModels(Car, "1")

		if err != nil && response == nil {
			t.Errorf("Shouldn't have an error, got: %s", err.Error())
		}
	})

	t.Run("Should return error when response is invalid", func(t *testing.T) {
		expected := "Parâmetros inválidos"
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarModelos", mockHandler(mockError))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrandModels(Car, "1")

		if err == nil && response != nil {
			t.Errorf("Should have an error, got response: %s", response.Modelos[0].Label)
		}
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}

func TestFipeService_GetBrandModelYear(t *testing.T) {
	t.Run("Should return All Years from Models when the brand id and model id is valid", func(t *testing.T) {
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarAnoModelo", mockHandler(mockModelYears))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrandModelsYears(Car, "1", "1")

		if err != nil && response == nil {
			t.Errorf("Shouldn't have an error, got: %s", err.Error())
		}
	})

	t.Run("Should return error when response is invalid", func(t *testing.T) {
		expected := "Parâmetros inválidos"
		mux := http.NewServeMux()
		mux.HandleFunc("/api/veiculos/ConsultarTabelaDeReferencia", mockHandler(mockReferences))
		mux.HandleFunc("/api/veiculos/ConsultarAnoModelo", mockHandler(mockError))
		ts := httptest.NewServer(mux)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetBrandModelsYears(Car, "1", "1")

		if err == nil && response != nil {
			t.Errorf("Should have an error, got response: %s", response[0].Label)
		}
		if err.Error() != expected {
			t.Errorf("Expected: %s, got: %s", expected, err.Error())
		}
	})
}
