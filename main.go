package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// Struct to store API configuration, specifically the OpenWeatherMap API key
type apiConfig struct {
	OpenWeatherMapApiKey string `json:"OpenWeatherMapApiKey"`
}

// Struct to store weather data, primarily city name and temperature in Celsius
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Celsius float64 `json:"temp"`
	} `json:"main"`
}

func main() {
	// Set handler for the "/weather/" endpoint
	http.HandleFunc("/weather/", weatherHandler)

	// Print a message indicating that the server is starting and running on Port 8000
	fmt.Println("Start Server on Port 800")

	// Start the HTTP server and handle errors if any
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// Handler for the "/weather/" endpoint
func weatherHandler(w http.ResponseWriter, r *http.Request) {
	// Get the city name from the URL
	city := strings.SplitN(r.URL.Path, "/", 3)[2]

	// Call the query function to get weather data
	data, err := query(city)
	if err != nil {
		// If an error occurs, provide an HTTP response with status 500 (Internal Server Error)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header for the response as JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Encode weather data into JSON format and write it to the response
	json.NewEncoder(w).Encode(data)
}

// Function to query the OpenWeatherMap API
func query(city string) (weatherData, error) {
	// Read API configuration from a file
	apiConfig, err := loadApiConfig(".apiConfig")
	if err != nil {
		return weatherData{}, err
	}

	// Perform an HTTP GET request to the OpenWeatherMap API
	r, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapApiKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	// Close the response body when done
	defer r.Body.Close()

	// Read and decode the JSON response into the weatherData struct
	var d weatherData
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	// Convert temperature from Kelvin to Celsius
	d.Main.Celsius = kelvinToCelcius(d.Main.Celsius)

	return d, nil
}

// Function to convert temperature from Kelvin to Celsius
func kelvinToCelcius(kelvin float64) float64 {
	return kelvin - 273.15
}

// Function to load API configuration from a file
func loadApiConfig(filename string) (apiConfig, error) {
	// Read the entire content of the file
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return apiConfig{}, err
	}

	// Create a variable to store the configuration data
	var c apiConfig

	// Decode JSON from the file content into the apiConfig struct
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfig{}, err
	}

	return c, nil
}
