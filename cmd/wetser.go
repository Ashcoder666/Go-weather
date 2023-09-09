/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"

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

type places struct {
	Name    string `json:"name"`
	Region  string `json:"region"`
	Country string `json:"Country"`
}

func locationSearch() {
	fmt.Println("Enter your location")
	var location string
	fmt.Scan(&location)
	fullUrl := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/search.json?q=%s", location)

	// if err != nil {
	// 	return nil, err
	// }

	defer resp.Body.Close()

	responses, _ := io.ReadAll(resp.Body)

	fmt.Println(string(responses))
	// var searchResponse []places

	// if err := json.NewDecoder(resp.Body).Decode(&searchResponse); err != nil {
	// 	return nil, err
	// }

	// return searchResponse, nil
}
