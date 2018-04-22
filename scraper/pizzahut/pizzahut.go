package pizzahut

import (
	"strings"

	"../../lib/json"
	"../../models"
)

type Scraper struct {
	Url      string
	Response *apiResponse
}

type storeResponse struct {
	Id                         int    `json:"id"`
	Code                       string `json:"code"`
	Name                       string `json:"name"`
	AddressLine1               string `json:"addressLine1"`
	AddressLine2               string `json:"addressLine2"`
	City                       string `json:"city"`
	State                      string `json:"state"`
	PostCode                   string `json:"postCode"`
	DeliveryTimeMinutes        int    `json:"deliveryTimeMinutes"`
	PickupTimeMinutes          int    `json:"pickupTimeMinutes"`
	KilometersTo               int    `json:"kilometersTo"`
	CanPickupNow               bool   `json:"canPickupNow"`
	CanAcceptOnlineOrders      bool   `json:"canAcceptOnlineOrders"`
	TimeBeforeStoreOpens       string `json:"timeBeforeStoreOpens"`
	IsOpen                     bool   `json:"isOpen"`
	Telephone                  string `json:"telephone"`
	TodayPickupHours           string `json:"todayPickupHours"`
	TodayDeliveryHours         string `json:"todayDeliveryHours"`
	TodayDineInHours           string `json:"todayDineInHours"`
	TradingStatus              string `json:"tradingStatus"`
	PhStoreUrl                 string `json:"phStoreUrl"`
	CanAcceptOnlineCreditCards bool   `json:"canAcceptOnlineCreditCards"`
	CanAcceptPayPalPayment     bool   `json:"canAcceptPayPalPayment"`
	CanVirtualDrivethru        bool   `json:"canVirtualDrivethru"`
	CanPickup                  bool   `json:"canPickup"`
	CanDeliver                 bool   `json:"canDeliver"`
	Location                   string `json:"location"`
	TimeTable                  string `json:"timeTable"`
}

type storesResponse struct {
	Stores []storeResponse `json:"stores"`
}

type apiResponse struct {
	Success bool           `json:"success"`
	Result  storesResponse `json:"result"`
}

func CreateScraper(url string) *Scraper {
	scraper := new(Scraper)
	scraper.Url = url
	scraper.Response = scraper.GetResponse()

	return scraper
}

func (scraper *Scraper) Scrape() []*data.StoreInfo {
	return scraper.ParseStores()
}

func (scraper *Scraper) GetResponse() *apiResponse {
	response := new(apiResponse)
	err := json.Get(scraper.Url, &response)
	if err != nil {
		panic(err.Error())
	}
	return response
}

func (scraper Scraper) GetName() string {
	return "PizzaHut"
}

func (scraper *Scraper) ParseStores() []*data.StoreInfo {
	var stores []*data.StoreInfo

	for _, store := range scraper.Response.Result.Stores {
		stores = append(stores, scraper.ParseStore(store))
	}
	return stores
}

func (scraper *Scraper) ParseStore(jsonStore storeResponse) *data.StoreInfo {
	location := strings.Split(jsonStore.Location, ",")
	store := new(data.StoreInfo)

	store.Name = jsonStore.Name
	store.Address = jsonStore.AddressLine1 + " " + jsonStore.AddressLine2
	store.Latitude = location[0]
	store.Longitude = location[1]

	return store
}
