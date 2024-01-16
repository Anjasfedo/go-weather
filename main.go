package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type apiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
}

func main() {

}

func loadApiConfig(filename string) (apiConfig, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfig{}, err
	}

	var c apiConfig

	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfig{}, err
	}

	return c, nil
}
