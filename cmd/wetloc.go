/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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
	// resp, err := http.Get("http://httpbin.org/ip")

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }

	// body, _ := io.ReadAll(resp.Body)

	// ip := string(body)
	// fmt.Println("Your IP address is:", ip)

	ipAddress, _ := getIP()

	// lat, lon, _ := getLatLonByIP(ipAddress)

	fullUrl := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/current.json?q=%s", ipAddress)

	GetWeatherDeatils(fullUrl, "GET", 6)

}

type WeatherResponse struct {
	Location struct {
		Name    string `json:"name"`
		Region  string `json:"region"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC float32 `json:"temp_c"`
	} `json:"current"`
}

type IPResponse struct {
	Origin string `json:"origin"`
}

func GetWeatherDeatils(url, method string, timeout int) {
	client := http.Client{Timeout: time.Duration(timeout) * time.Second}
	req, err := http.NewRequest(method, url, nil)
	// fmt.Println(req)
	if err != nil {

		fmt.Println(err)
	}

	headersMap := make(map[string]string)

	headersMap["x-rapidapi-host"] = "weatherapi-com.p.rapidapi.com"
	headersMap["x-rapidapi-key"] = "c256ecdf6emshcedc5242d337c13p19a862jsnee79af877651"
	// req.Header.Add(headers)

	for index, header := range headersMap {
		// fmt.Println(index, header)
		req.Header.Add(index, header)
	}

	res, err := client.Do(req)
	if err != nil {

		fmt.Println(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Println("breaaked")
	}

	fmt.Println(string(body))

	var weatherData WeatherResponse

	errr := json.Unmarshal(body, &weatherData)

	if errr != nil {
		fmt.Println("breaaked")
	}

	// fmt.Printf("Body : %s", body)
	fmt.Printf("Location: %s, %s, %s\n", weatherData.Location.Name, weatherData.Location.Region, weatherData.Location.Country)
	fmt.Printf("Temperature (Celsius): %.1f\n", weatherData.Current.TempC)
}

func getLatLonByIP(ipAddress string) (float64, float64, error) {
	url := "http://ip-api.com/json/" + ipAddress

	// Make an HTTP GET request to the IP geolocation API
	resp, err := http.Get(url)
	if err != nil {
		return 0.0, 0.0, err
	}
	defer resp.Body.Close()

	var data struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, 0, err
	}

	return data.Lat, data.Lon, nil
}

func getIP() (string, error) {
	// Make an HTTP GET request to httpbin.org to fetch your IP address
	resp, err := http.Get("http://httpbin.org/ip")
	if err != nil {
		return "", errors.New("An error occurred: " + err.Error())
	}
	defer resp.Body.Close()

	var ipResponse IPResponse
	if err := json.NewDecoder(resp.Body).Decode(&ipResponse); err != nil {
		return "", err
	}
	// Read the response body
	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	return "", errors.New("An error occurred: " + err.Error())
	// }

	return ipResponse.Origin, nil
}
