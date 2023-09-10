/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// wetserCmd represents the wetser command
var wetserCmd = &cobra.Command{
	Use:   "wetser",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		locationSearch()
		// GetWeatherDeatils(fullUrl, "GET", 6)
	},
}

func init() {
	rootCmd.AddCommand(wetserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wetserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wetserCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Places struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

func locationSearch() {
	fmt.Println("Enter your location")
	var location string
	fmt.Scan(&location)
	client := http.Client{Timeout: time.Duration(5) * time.Second}
	fullUrl := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/search.json?q=%s", location)
	req, err := http.NewRequest("GET", fullUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "c256ecdf6emshcedc5242d337c13p19a862jsnee79af877651")

	res, err := client.Do(req)
	if err != nil {

		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("breaaked")
	}

	// fmt.Println(string(body))

	var placesInstance []Places

	errr := json.Unmarshal(body, &placesInstance)

	if errr != nil {
		fmt.Println(errr)
	} else {
		// Now, placesInstance contains the data from the JSON response
		for _, place := range placesInstance {
			fmt.Println("Name:", place.Name)
			fmt.Println("Region:", place.Region)
			fmt.Println("Country:", place.Country)
			fmt.Println("---")
		}
	}

}
