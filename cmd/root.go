/*
Copyright © 2023 CRISTÓVÃO OLEGÁRIO DE CASTRO <CRISTOVAOOLEGARIO@GMAIL.COM>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "fipe-scrapper",
	Short: "CLI tool to gather data from fipe-api -> https://veiculos.fipe.org.br",
	Long: `CLI tool to gather data from fipe-api -> https://veiculos.fipe.org.br .
	The Fipe Table expresses average vehicle prices in the national market, 
	serving only as a parameter for the process or estimate. 
	Prices differ depending on the region, conservation, color, 
	accessories or any other factor that may influence supply and demand conditions for a specific vehicle.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.fipe-scrapper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
