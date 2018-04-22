package main

import (
	"./lib/webscraper"
	"./models"
	"./stores/dominos"
	"./stores/pizzahut"
)

const (
	csvDir string = "./csv/"
)

func main() {
	storeScrapers := make([]data.StoreScraper, 2)
	storeScrapers[0] = pizzahut.StoreScraper{}
	storeScrapers[1] = dominos.StoreScraper{}

	for _, store := range storeScrapers {
		webscraper.Scrape(store, csvDir+store.GetName()+".csv")
	}

}
