package main

import (
	"fmt"
	"strconv"

	"./lib/storeRepo"
	"./models"
	"./scraper/dominos"
	"./scraper/pizzahut"
)

const (
	csvDir string = "./csv/"
)

func main() {

	scrapers := getScrapers()
	channels := createChannels(scrapers)

	for index, scraper := range scrapers {
		go scrape(scraper, channels[index])
	}

	waitResponse(channels)

}

func scrape(scraper data.StoreScraper, done chan bool) {
	branches := scraper.Scrape()
	storeRepo.SaveToCSV(branches, csvDir+scraper.GetName()+".csv")

	done <- true
}

func waitResponse(channels []chan bool) {
	for index, done := range channels {
		fmt.Println("task #" + strconv.Itoa(index+1) + " done? = " + strconv.FormatBool(<-done))
	}
}

func createChannels(scrapers []data.StoreScraper) []chan bool {
	channels := make([]chan bool, len(scrapers))
	for i := 0; i < len(scrapers); i++ {
		channels[i] = make(chan bool)
	}
	return channels
}

func getScrapers() []data.StoreScraper {

	domino := dominos.CreateScraper("https://www.dominos.com.au/store-finder")
	//due to time constraints, haven't done a crawler to click and wait on ajax results, therefore have used directly the api behind the ajax calls
	pizzahut := pizzahut.CreateScraper("https://ecommerce-deploy-61-api.prod.pizzahutaustralia.com.au/api/services/app/store/GetTradingStoresStates")

	return []data.StoreScraper{pizzahut, domino}
}
