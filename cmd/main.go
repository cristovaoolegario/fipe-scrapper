package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/services"
)

func main() {
	service := &services.FipeService{
		Client: &http.Client{
			Timeout: time.Duration(10) * time.Second,
		},
	}
	dto, err := service.GetBrands(services.Car)
	if err != nil {
		panic(err)
	}

	for _, item := range dto {
		fmt.Printf("%s - %s\n", item.Label, item.Value)
	}
}
