/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/services"
	"github.com/spf13/cobra"
)
var vehicleType int16

// brandsCmd represents the brands command
var brandsCmd = &cobra.Command{
	Use:   "brands",
	Short: "List all brands",
	RunE: func(cmd *cobra.Command, _ []string) error {
		if vehicleType <= 0 || vehicleType > 2{
			return errors.New("invalid vehicle type")
		}
		
		service := services.ProvideDefaultService()
		brands, err := service.GetBrands(services.Vehicle(vehicleType))

		if err != nil {
			return err
		}

		fmt.Printf("%v", brands)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(brandsCmd)
	brandsCmd.Flags().Int16VarP(&vehicleType, "type","t", 1, "type of vehicle (Car = 1 | Bike = 2)")
	brandsCmd.MarkFlagRequired("type")
}
