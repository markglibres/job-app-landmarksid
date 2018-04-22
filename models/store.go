package data

type StoreInfo struct {
	Name, Address       string
	Latitude, Longitude string
}

type StoreScraper interface {
	GetBranchStores() []StoreInfo
	GetName() string
}
