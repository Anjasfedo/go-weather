package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiConfig struct{
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

type weatherData struct{
	Name string `json:"name"`
	Main struct{
		Kelvin float64 `json:"temp"`
	}`json:"main"`
}

func main() {

}
