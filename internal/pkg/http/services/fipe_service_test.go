package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const mockReferences = `[{"Codigo": 295,"Mes": "março/2023 "},{"Codigo": 294,"Mes": "fevereiro/2023 "}]`

func mockHandler(mockData string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := strings.NewReader(mockReferences)
		byteValue, _ := ioutil.ReadAll(mockResponse)

		w.Header().Add("Content-Type", "application/json")
		w.Write(byteValue)
	})
}

func TestFipeService_GetAllReferences(t *testing.T) {
	t.Run("Should return References data when response is valid", func(t *testing.T) {
		ts := httptest.NewServer(mockHandler(mockReferences))

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetAllReferences()

		if err != nil && response == nil {
			t.Errorf("Shouldn't have an error, got: %s", err)
		}
	})
}

func TestFipeService_GetLatestReference(t *testing.T) {
	t.Run("Should return the latest reference code when response is valid", func(t *testing.T) {
		ts := httptest.NewServer(mockHandler(mockReferences))

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
}
