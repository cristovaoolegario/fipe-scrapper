package services

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const mockReferences = `[{"Codigo": 295,"Mes": "março/2023 "},{"Codigo": 294,"Mes": "fevereiro/2023 "}]`

func TestFipeService_GetAllReferences(t *testing.T) {
	t.Run("Should return References data when response is valid", func(t *testing.T) {
		ts := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mockResponse := strings.NewReader(mockReferences)
				byteValue, _ := ioutil.ReadAll(mockResponse)

				w.Header().Add("Content-Type", "application/json")
				w.Write(byteValue)
			}),
		)

		defer ts.Close()

		service := NewFipeService(http.Client{}, ts.URL)

		response, err := service.GetAllReferences()

		if err != nil && response == nil {
			t.Errorf("Error: %s", err)
		}
	})
}
