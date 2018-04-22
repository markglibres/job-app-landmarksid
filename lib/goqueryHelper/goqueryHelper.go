package goqueryHelper

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetDocument(url string) *goquery.Document {
	response := GetResponse(url)
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		panic(err.Error())
	}
	return doc
}

func GetResponse(url string) *http.Response {
	response, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}
	return response
}
