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

	references, _ := service.GetAllReferences()

	fmt.Println("\n====== Reference table ======")

	for _, item := range references {
		fmt.Printf("%s - %d\n", item.Mes, item.Codigo)
	}

	brandsDto, _ := service.GetBrands(services.Car)

	fmt.Println("\n====== Brands table ======")

	for _, item := range brandsDto {
		fmt.Printf("%s - %s\n", item.Label, item.Value)
	}

	modelsBrands, _ := service.GetBrandModels(services.Car, brandsDto[0].Value)

	fmt.Println("\n====== Models table ======")

	for _, item := range modelsBrands.Modelos {
		fmt.Printf("%s - %d\n", item.Label, item.Value)

	}
}
