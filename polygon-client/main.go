package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

func main() {

	// avoid hardcoding api key for codebase
	APIKeyLocation := "C:\\Users\\docog\\OneDrive\\Documents\\polygon_api_key.txt"

	setAPIKey(APIKeyLocation)

	// creates a new client using API key
	c := polygon.New(os.Getenv("POLYGON_API_KEY"))

	// setting the parameters for API requests
	params := models.GetDailyOpenCloseAggParams{
		Ticker: "AAPL",
		Date:   models.Date(time.Date(2023, 3, 8, 0, 0, 0, 0, time.Local)),
	}.WithAdjusted(true)

	// makes the request using the parameters
	apiResquest, err := c.GetDailyOpenCloseAgg(context.Background(), params)
	if err != nil {
		log.Fatal("Resulted in: ", err)
	}

	fmt.Println(apiResquest) // do something with the result
}

func setAPIKey(fileName string) {
	file, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	apiKey := string(file)

	os.Setenv("POLYGON_API_KEY", apiKey)
}
