package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/services"
)

func main() {
	service := services.NewFipeService(http.Client{
		Timeout: time.Duration(10) * time.Second,
	}, services.BASE_URL)

	brandsDto, _ := service.GetAllReferences()

	for _, item := range brandsDto {
		fmt.Printf("%s - %d", item.Mes, item.Codigo)
	}

	dto, _ := service.GetBrands(services.Car)

	for _, item := range dto {
		fmt.Printf("%s - %s\n", item.Label, item.Value)
	}
}
