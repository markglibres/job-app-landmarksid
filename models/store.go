package data

type StoreInfo struct {
	Name, Address       string
	Latitude, Longitude string
}

type StoreScraper interface {
	Scrape() []*StoreInfo
	GetName() string
}

type ChannelInfo struct {
	Name   string
	IsDone bool
}
