package data

type StoreInfo struct {
	Name, Address       string
	Latitude, Longitude float64
}

type StoreScraper interface {
	GetBranchStores() []StoreInfo
	GetName() string
}
