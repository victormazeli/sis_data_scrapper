package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Struct for API response (change based on your API's structure)
type APIResponse struct {
	Data       []interface{} `json:"data"`       // Store actual data here
	TotalPages int           `json:"totalPages"` // Example pagination field
	Page       int           `json:"page"`       // Current page
}

// FetchData from API (Handles pagination and rate limiting)
func fetchData(apiURL string, page int, rateLimit time.Duration, ch chan []interface{}) {
	for {
		urlWithParams := fmt.Sprintf("%s?page=%d", apiURL, page)
		resp, err := http.Get(urlWithParams)

		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}

		if resp.StatusCode != 200 {
			log.Fatalf("Non-OK HTTP status: %d", resp.StatusCode)
		}

		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Parse JSON response
		var apiResponse APIResponse
		err = json.Unmarshal(body, &apiResponse)
		if err != nil {
			log.Fatalf("Error unmarshaling JSON: %v", err)
		}

		// Send the fetched data to the channel for processing
		ch <- apiResponse.Data

		// Break if last page is reached
		if apiResponse.Page >= apiResponse.TotalPages {
			close(ch)
			return
		}

		// Increment page for the next iteration
		page++

		// Respect API rate limits
		time.Sleep(rateLimit)
	}
}

// SaveData writes fetched data to JSON file incrementally
func saveDataToFile(fileName string, ch chan []interface{}) {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Start writing data incrementally
	for data := range ch {
		// Convert data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatalf("Error marshaling data: %v", err)
		}

		// Write JSON data to file
		_, err = file.Write(jsonData)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}

		// Add a newline for readability
		file.WriteString("\n")
	}
}

func main() {
	// Parse CLI arguments
	apiURL := flag.String("url", "", "API endpoint URL")
	outputFile := flag.String("out", "output.json", "Output JSON file")
	startPage := flag.Int("start", 1, "Starting page for scraping")
	rateLimit := flag.Int("rate", 1, "Rate limit in seconds between requests")

	flag.Parse()

	if *apiURL == "" {
		log.Fatalf("API URL is required")
	}

	// Channel for fetching data
	dataChannel := make(chan []interface{}, 1)

	// Start fetching data
	go fetchData(*apiURL, *startPage, time.Duration(*rateLimit)*time.Second, dataChannel)

	// Start saving data to file
	saveDataToFile(*outputFile, dataChannel)

	fmt.Println("Data scraping completed and saved to", *outputFile)
}
