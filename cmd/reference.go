/*
Copyright © 2023 CRISTÓVÃO OLEGÁRIO

*/
package cmd

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cristovaoolegario/fipe-scrapper/internal/pkg/http/services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// referenceCmd represents the reference command
var referenceCmd = &cobra.Command{
	Use:   "reference",
	Short: "Update reference in config file",
	Long:  ``,
	Run: func(_ *cobra.Command, _ []string) {
		service := services.NewFipeService(http.Client{
			Timeout: time.Duration(10) * time.Second,
		}, services.BASE_URL)

		ref, err := service.GetLatestReference()

		if err != nil {
			fmt.Println(err)
			return
		}

		viper.Set("reference", ref)
		err = viper.WriteConfig()

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\nReference key updated to: %s\n", ref)
	},
}

func init() {
	rootCmd.AddCommand(referenceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// referenceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// referenceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
