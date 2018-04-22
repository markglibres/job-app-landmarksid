package main

import (
	"./lib"
	"./models"
	"./stores"
)

const (
	csvDir string = "./csv/"
)

func main() {
	storeScrapers := make([]data.StoreScraper, 2)
	storeScrapers[0] = stores.PizzaHut{}
	storeScrapers[1] = stores.Dominos{}

	scrapeDone := make(chan bool, len(storeScrapers))

	for _, store := range storeScrapers {
		go func(store data.StoreScraper, fullPath string) {
			branches := webscraper.Scrape(store)
			webscraper.SaveToCSV(branches, fullPath)
			scrapeDone <- true
		}(store, csvDir+store.GetName()+".csv")

	}

	<-scrapeDone
	close(scrapeDone)

}
