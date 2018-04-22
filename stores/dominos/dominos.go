package dominos

import (
	"../../models"
)

type StoreScraper struct {
}

func (store StoreScraper) GetBranchStores() []data.StoreInfo {
	branches := make([]data.StoreInfo, 1)
	branches[0] = data.StoreInfo{Name: "Dominos 1", Latitude: "-35.018596", Longitude: "117.883642"}
	return branches
}

func (store StoreScraper) GetName() string {
	return "Dominos"
}
