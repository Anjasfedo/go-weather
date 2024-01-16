# Go Weather API with OpenWeatherMap 🚀

This Go program provides a simple API to fetch weather data using the OpenWeatherMap API. It includes functionality to convert temperature from Kelvin to Celsius.

## Usage 🛠️

1. Make sure you have an API key from OpenWeatherMap. Create a file named .apiConfig in the project root and add your API key in the following format:

    ```.apiConfig
     {
        "OpenWeatherMapApiKey": "your_api_key"
     }
    ```

2. run the program:

    ```
     go run main.go
    ```

The server will start at `http://localhost:8000`.

## API Endpoints 🚪

### Get Weather by City

-   Endpoint: GET /weather/:city
-   Description: Retrieve weather data for a specific city.
-   Example: GET `http://localhost:8000/weather/bengkulu`
    The response includes the city name and temperature in Celsius.

## Code Explanation 📜

### main Function 🚀

#### Set Handler for Endpoint:

-   Sets up a handler for the /weather/:city endpoint.

#### Print Server Start Message:

-   Prints a message indicating that the server is starting on Port 8000.

#### Start HTTP Server:

-   Starts the HTTP server and handles any errors that may occur.

### weatherHandler Function 🚪

#### Extract City from URL:

-   Parses the city name from the URL.

#### Query OpenWeatherMap API:

-   Calls the query function to retrieve weather data for the specified city.

#### Set Response Header:

-   Sets the Content-Type header for the response as JSON.

#### Encode and Write Response:

-   Encodes the weather data into JSON format and writes it to the response.

### query Function 📊

#### Load API Configuration:

-   Reads the OpenWeatherMap API key from the .apiConfig file.

#### Perform API Request:

-   Sends an HTTP GET request to the OpenWeatherMap API with the specified city.

#### Close Response Body:

-   Ensures the response body is closed when done.

#### Decode JSON Response:

-   Decodes the JSON response from the API into the weatherData struct.

#### Convert Temperature:

-   Converts the temperature from Kelvin to Celsius.

#### Return Weather Data:

-   Returns the weather data.

### loadApiConfig Function 🌐

#### Read and Decode .apiConfig File:

-   Reads the content of the .apiConfig file and decodes it into the apiConfig struct.

#### Return API Configuration:

-   Returns the API configuration.

## Closing Notes 📝

Feel free to adjust the configuration, and if you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

Happy coding! 🚀👨‍💻
