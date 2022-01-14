package cmd

import (
	"fmt"
	"os"

	"github.com/aksharau/GoGormExamples/pkg/db"
	"github.com/aksharau/GoGormExamples/pkg/model"
	"github.com/aksharau/GoGormExamples/pkg/rest"
	"github.com/spf13/cobra"
)

var City string
var ApiKey string

var rootCmd = &cobra.Command{
	Use:   "weather [City] [ApiKey]",
	Short: "Gives city weather",
	Long:  `Gives city weather and also changes it ;-)`,

	Run: func(cmd *cobra.Command, args []string) {
		restC := rest.GetRestClient(ApiKey)
		weather := restC.GetWeather(City)
		dbRec := model.MapToCityWeather(weather)
		fmt.Println("REc is %+v", dbRec)
		dbC := db.GetDBClient()

		dbC.SaveWeather(dbRec)

		for _, rec := range dbC.GetAllRec() {
			fmt.Println("The data is %+v", rec)
		}

	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&City, "city", "c", "", "city")
	rootCmd.MarkFlagRequired("city")
	rootCmd.PersistentFlags().StringVarP(&ApiKey, "apiKey", "a", "", "open weather API key")
	rootCmd.MarkFlagRequired("apiKey")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
