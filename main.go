package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
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

type Urls struct {
	Url        string
	StartPage  int
	OutputFile string
}

// FetchData from API (Handles pagination and rate limiting)
func fetchData(apiURL string, page int, rateLimit time.Duration, ch chan []interface{}, cookie1 string, cookie2 string) {
	for {
		urlWithParams := fmt.Sprintf("%s?start=%d&max=1000", apiURL, page)
		response, err := http.NewRequest("GET", urlWithParams, nil)
		//resp, err := http.Get(urlWithParams)
		if err != nil {
			log.Fatalf("Error fetching data: %v", err)
		}
		cookieOne := &http.Cookie{Name: "JSESSIONID", Value: cookie1}
		cookieTwo := &http.Cookie{Name: "cas-session", Value: cookie2}

		response.AddCookie(cookieOne)
		response.AddCookie(cookieTwo)

		client := &http.Client{}

		resp, err := client.Do(response)

		//if resp.StatusCode != 200 {
		//	log.Println("Error is this ", err)
		//	log.Println("Response is ", resp.Request)
		//	log.Fatalf("Non-OK HTTP status: %d and body is %s", resp.StatusCode, resp.Body)
		//}

		defer resp.Body.Close() // Ensure the response body is closed

		if resp.StatusCode != http.StatusOK {
			bodyBytes, err := io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("Error reading body: %v", err)
			}
			bodyString := string(bodyBytes)
			log.Printf("Non-OK HTTP status: %d body is %s and apiurl is %s ", resp.StatusCode, bodyString, apiURL)
		}

		defer resp.Body.Close()

		// Read response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Error reading response body: %v", err)
		}

		// Parse JSON response
		var apiResponse APIResponse
		err = json.Unmarshal(body, &apiResponse.Data)
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
	//apiURL := flag.String("url", "", "API endpoint URL")
	//outputFile := flag.String("out", "output.json", "Output JSON file")
	//startPage := flag.Int("start", 1, "Starting page for scraping")
	jsessionId := flag.String("JSESSIONID", "", "Jsession Id")
	casSession := flag.String("cas-session", "", "cas-session Id")
	rateLimit := flag.Int("rate", 5, "Rate limit in seconds between requests")

	flag.Parse()
	if jsessionId == nil {
		log.Fatalf("JsessionId is required")
	}
	if casSession == nil {
		log.Fatalf("Cas session is required")
	}

	for _, value := range getUrl() {
		dataChannel := make(chan []interface{}, 1)

		go fetchData(value.Url, value.StartPage, time.Duration(*rateLimit)*time.Second, dataChannel, *jsessionId, *casSession)

		saveDataToFile(value.OutputFile, dataChannel)
	}
	// Start saving data to file

	fmt.Println("Data scraping completed and saved to")
}

func getUrl() []Urls {
	return []Urls{
		{
			Url:        "https://portal.miva.university/enrol/rest/application",
			StartPage:  0,
			OutputFile: "applications.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/bundle",
			StartPage:  0,
			OutputFile: "bundles.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/agreement",
			StartPage:  0,
			OutputFile: "agreements.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/assessment",
			StartPage:  0,
			OutputFile: "assessments.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/bank_account",
			StartPage:  0,
			OutputFile: "bankAccounts.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/building",
			StartPage:  0,
			OutputFile: "buildings.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/enrolment/bundle",
			StartPage:  0,
			OutputFile: "enrollmentBundle.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/comment",
			StartPage:  0,
			OutputFile: "comment.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/contact",
			StartPage:  0,
			OutputFile: "contacts.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/enrolment/course",
			StartPage:  0,
			OutputFile: "courseEnrollments.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/courseOffering",
			StartPage:  0,
			OutputFile: "courseOfferings.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/course-result-template",
			StartPage:  0,
			OutputFile: "courseResultTemplates.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/course",
			StartPage:  0,
			OutputFile: "courses.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/debtoraccount",
		//	StartPage:  0,
		//	OutputFile: "debtorAccounts.json",
		//},
		// Debtor is having status no content
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/debtor",
		//	StartPage:  0,
		//	OutputFile: "debtors.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/document",
			StartPage:  0,
			OutputFile: "documents.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/employee",
		//	StartPage:  0,
		//	OutputFile: "employees.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/employmentrecord",
			StartPage:  0,
			OutputFile: "employmentRecords.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/event",
			StartPage:  0,
			OutputFile: "events.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/financialAid",
			StartPage:  0,
			OutputFile: "financialAids.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/financialApplication",
			StartPage:  0,
			OutputFile: "financialApplications.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/grade-scale",
			StartPage:  0,
			OutputFile: "gradeScales.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/image",
			StartPage:  0,
			OutputFile: "images.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/institution",
			StartPage:  0,
			OutputFile: "institutions.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/intake",
			StartPage:  0,
			OutputFile: "intakes.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/invoice",
			StartPage:  0,
			OutputFile: "invoices.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/lead",
		//	StartPage:  0,
		//	OutputFile: "leads.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/order",
			StartPage:  0,
			OutputFile: "orders.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/org",
			StartPage:  0,
			OutputFile: "organizations.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/person/affiliation",
		//	StartPage:  0,
		//	OutputFile: "personAffiliations.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/productFee",
			StartPage:  0,
			OutputFile: "productFees.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/product",
			StartPage:  0,
			OutputFile: "products.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/enrolment/program-level",
			StartPage:  0,
			OutputFile: "programLevelEnrollments.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/program",
			StartPage:  0,
			OutputFile: "programs.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/prospect",
			StartPage:  0,
			OutputFile: "prospects.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/receipt",
			StartPage:  0,
			OutputFile: "receipts.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/referral",
		//	StartPage:  0,
		//	OutputFile: "referrals.json",
		//},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/enrolment/registration",
		//	StartPage:  0,
		//	OutputFile: "enrolmentRegistration.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/report",
			StartPage:  0,
			OutputFile: "reports.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/schoolSubject",
			StartPage:  0,
			OutputFile: "schoolSubject.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/enrolment/short-course",
			StartPage:  0,
			OutputFile: "shortCourseEnrolments.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/state",
			StartPage:  0,
			OutputFile: "states.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/account",
			StartPage:  0,
			OutputFile: "studentAccounts.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/student",
			StartPage:  0,
			OutputFile: "students.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/timetable",
			StartPage:  0,
			OutputFile: "timeTables.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/type",
		//	StartPage:  0,
		//	OutputFile: "types.json",
		//},
		{
			Url:        "https://portal.miva.university/enrol/rest/voucher",
			StartPage:  0,
			OutputFile: "vouchers.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/webhook",
			StartPage:  0,
			OutputFile: "webhooks.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/authorization",
			StartPage:  0,
			OutputFile: "users.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/population",
			StartPage:  0,
			OutputFile: "populations.json",
		},
		{
			Url:        "https://portal.miva.university/enrol/rest/authorization/role",
			StartPage:  0,
			OutputFile: "role.json",
		},
		//{
		//	Url:        "https://portal.miva.university/enrol/rest/statistics/ids",
		//	StartPage:  0,
		//	OutputFile: "statistics.json",
		//},
	}
}
