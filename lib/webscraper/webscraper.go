package webscraper

import (
	"encoding/csv"
	"log"
	"os"

	"../../models"
)

var fileHeaders = []string{"Name", "Address", "Latitude", "Longitude"}

func Scrape(store data.StoreScraper, fullPath string) {

	results := store.GetBranchStores()
	SaveToCSV(results, fullPath)
}

func SaveToCSV(branches []data.StoreInfo, filePath string) {

	file, err := os.Create(filePath)
	checkError(err, "Cannot create file")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(fileHeaders)
	checkError(err, "Error writing to file")

	for _, branch := range branches {
		writer.Write([]string{branch.Name, branch.Address, branch.Latitude, branch.Longitude})
	}

}

func checkError(err error, message string) {
	if err != nil {
		log.Fatal(message, err)
	}
}
