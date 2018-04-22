package pizzahut

import (
	"strings"

	"../../lib/jsonService"
	"../../models"
)

const (
	storesApiEndPoint string = "https://ecommerce-deploy-61-api.prod.pizzahutaustralia.com.au/api/services/app/store/GetTradingStoresStates"
)

type StoreScraper struct {
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

func (store StoreScraper) GetBranchStores() []data.StoreInfo {
	jsonStores := getStoresFromApi(storesApiEndPoint)
	return parseStores(jsonStores)
}

func (store StoreScraper) GetName() string {
	return "PizzaHut"
}

func getStoresFromApi(url string) []storeResponse {
	result := apiResponse{}
	err := jsonService.GetJson(url, &result)
	if err != nil {
		panic(err.Error())
	}
	return result.Result.Stores
}

func parseStores(jsonStores []storeResponse) []data.StoreInfo {
	stores := make([]data.StoreInfo, len(jsonStores))
	for index, store := range jsonStores {
		stores[index] = parseStore(store)
	}
	return stores
}

func parseStore(jsonStore storeResponse) data.StoreInfo {
	location := strings.Split(jsonStore.Location, ",")
	return data.StoreInfo{
		Name:      jsonStore.Name,
		Address:   jsonStore.AddressLine1 + " " + jsonStore.AddressLine2,
		Latitude:  location[0],
		Longitude: location[1]}
}
