package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// makes request to the weather api and handles the response
func MakeRequest(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("request failed with status code: %v", response.StatusCode)

	}

	log.Printf("Request: %v", response.Status)

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

// Main request handler called on / endpoint
func QueryHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("%v %v\n", request.Method, request.URL)

	// Set the response header to return a json
	responseWriter.Header().Set("Content-Type", "application/json")

	// Check if the city is passed in request query
	if !request.URL.Query().Has("city") {
		io.WriteString(responseWriter, "'city' in query is required")
		return
	}

	// get the city
	city := request.URL.Query().Get("city")
	log.Printf("City: %v\n", city)

	// Get the API_KEY
	ENV_CONST, _ := godotenv.Read()

	API_KEY := ENV_CONST["API_KEY"]

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%v&q=%v&aqi=no", API_KEY, city)

	// Make request to get the weather data
	response, err := MakeRequest(url)

	if err != nil {
		io.WriteString(responseWriter, err.Error())
		return
	}

	// write the JSON to the response writer
	io.WriteString(responseWriter, string(response))
}

// Main request handler called on / endpoint
func FormatQueryHandler(responseWriter http.ResponseWriter, request *http.Request) {
	log.Printf("%v %v\n", request.Method, request.URL)

	// Set the response header to return a json
	responseWriter.Header().Set("Content-Type", "application/json")

	// Check if the city is passed in request query
	if !request.URL.Query().Has("city") {
		io.WriteString(responseWriter, "'city' in query is required")
		return
	}

	// Get the city
	city := request.URL.Query().Get("city")
	log.Printf("City: %v\n", city)

	// Get the API_KEY
	ENV_CONST, _ := godotenv.Read()
	API_KEY := ENV_CONST["API_KEY"]

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%v&q=%v&aqi=no", API_KEY, city)

	/* Refactor below into a function */
	// Make request to get the weather data
	response, err := MakeRequest(url)

	if err != nil {
		io.WriteString(responseWriter, fmt.Sprintf(`{ "message": %v }`, err.Error()))
		return
	}

	// We could just return the []byte response from the make request
	// However, we will use the struct to remove some field
	// Basically, we are customizing the response

	/* ends here */

	/* refactor the below into a function */
	// create a variable for the Weather response
	var weatherResponse WeatherResponse

	// use json unmarshal and struct to bind the body (from MakeRequest)
	if err := json.Unmarshal(response, &weatherResponse); err != nil {
		io.WriteString(responseWriter, fmt.Sprintf(`{ "message": %v }`, err.Error()))
		return
	}

	// encode the weatherResponse struct to JSON
	weatherResponseJSON, err := json.Marshal(weatherResponse)

	if err != nil {
		io.WriteString(responseWriter, fmt.Sprintf(`{ "message": %v }`, err.Error()))
		return
	}
	/* ends here */

	// write the JSON to the response writer
	io.WriteString(responseWriter, string(weatherResponseJSON))
}
