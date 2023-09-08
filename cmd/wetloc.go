/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// wetlocCmd represents the wetloc command
var wetlocCmd = &cobra.Command{
	Use:   "wetloc",
	Short: "wetloc is used to get weather of specific location",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		getWeather()
	},
}

func init() {
	rootCmd.AddCommand(wetlocCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// wetlocCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// wetlocCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getWeather() {
	client := http.Client{Timeout: time.Duration(5) * time.Second}
	url := "https://weatherapi-com.p.rapidapi.com/current.json?q=kochi"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error : ", err)
	}

	req.Header.Add("x-rapidapi-host", "weatherapi-com.p.rapidapi.com")

	req.Header.Add("x-rapidapi-key", "c256ecdf6emshcedc5242d337c13p19a862jsnee79af877651")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error : ", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	fmt.Printf("Body : %s", body)
}
