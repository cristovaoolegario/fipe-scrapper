package health

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/services"
)

func CheckFipeApiHealth() {
	var (
		brandCode string
		modelCode string
	)
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

	brandCode = brandsDto[0].Value
	modelsBrands, _ := service.GetBrandModels(services.Car, brandCode)

	fmt.Println("\n====== Models table ======")

	for _, item := range modelsBrands.Modelos {
		fmt.Printf("%s - %d\n", item.Label, item.Value)

	}

	fmt.Println("\n====== Models Years table ======")

	modelCode = fmt.Sprint(modelsBrands.Modelos[0].Value)
	modelYears, _ := service.GetBrandModelsYears(services.Car, brandCode, modelCode)
	modelYearArr := strings.Split(modelYears[0].Value, "-")
	year, fuelType := modelYearArr[0], modelYearArr[1]

	for _, item := range modelYears {
		arr := strings.Split(item.Value, "-")
		year := arr[0]
		fuelType := arr[1]
		fmt.Printf("Year:%s Fuel type: %s\n", year, fuelType)
	}

	fmt.Println("\n====== Vehicles table ======")

	vehicle, _ := service.GetVehicle(services.Car, brandCode, modelCode, year, fuelType)

	fmt.Printf("%+v", vehicle)
}
