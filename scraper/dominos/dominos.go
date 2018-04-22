package dominos

import (
	"strings"

	"../../lib/goqueryHelper"
	"../../lib/urlHelper"
	"../../models"
	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	Url      string
	Document *goquery.Document
}

func CreateScraper(url string) *Scraper {
	scraper := new(Scraper)
	scraper.Url = url
	scraper.Document = goqueryHelper.GetDocument(url)
	return scraper
}

func (scraper *Scraper) Scrape() []*data.StoreInfo {

	var branches []*data.StoreInfo

	regions := scraper.GetRegionsUrl()
	var stores []string

	for _, region := range regions {
		regionScraper := CreateScraper(region)
		stores = append(stores, regionScraper.GetStoresUrlByRegion()...)
	}

	for _, store := range stores {
		storeScraper := CreateScraper(store)
		branches = append(branches, storeScraper.GetStoreInfo())
	}
	return branches
}

func (scraper *Scraper) GetName() string {
	return "Dominos"
}

func (scraper *Scraper) GetRegionsUrl() []string {
	doc := scraper.Document
	var regions []string
	doc.Find(".region-link > a").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			regions = append(regions, scraper.GetFullUrl(val))
		}
	})

	return regions
}

func (scraper *Scraper) GetStoresUrlByRegion() []string {
	doc := scraper.Document
	var stores []string
	doc.Find(".store-information > h4 > a").Each(func(i int, s *goquery.Selection) {
		val, exists := s.Attr("href")
		if exists {
			stores = append(stores, scraper.GetFullUrl(val))
		}
	})
	return stores
}

func (scraper *Scraper) GetFullUrl(path string) string {
	//todo: check if path is full url or not
	return urlHelper.Authority(scraper.Url) + path
}

func (scraper *Scraper) GetStoreInfo() *data.StoreInfo {
	doc := scraper.Document
	storeDoc := doc.Find(".info-row[aria-label='Store adress']").First()
	return ParseStoreInfo(storeDoc)
}

func ParseStoreInfo(storeDoc *goquery.Selection) *data.StoreInfo {

	store := new(data.StoreInfo)

	lat, exists := storeDoc.ChildrenFiltered("[name='store-lat']").Attr("value")
	if store.Latitude = ""; exists {
		store.Latitude = lat
	}
	long, exists := storeDoc.ChildrenFiltered("[name='store-lon']").Attr("value")
	if store.Longitude = ""; exists {
		store.Longitude = long
	}

	store.Name = storeDoc.Find(".store-details-text > .storetitle").Text()
	store.Address = strings.TrimSpace(storeDoc.Find("a").Text())
	return store
}
