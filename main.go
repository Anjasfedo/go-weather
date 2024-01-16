package main

import (
	"encoding/json"
	"fmt"
	"log"
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
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	http.HandleFunc("/weather/", weatherHandler)

	fmt.Println("Start Server on Port 800")

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Go!\n"))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	data, err := query(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application-json; charset=utf-8")
	json.NewEncoder(w).Encode(data)
}

func query(city string) (weatherData, error) {
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	r, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer r.Body.Close()

	var d weatherData
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	d.Main.Celsius = kelvinToCelcius(d.Main.Celsius)

	return d, nil
}

func kelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
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
