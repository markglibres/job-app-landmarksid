package stores

import (
	"../models"
)

type Dominos struct {
}

func (store Dominos) GetBranchStores() []data.StoreInfo {
	branches := make([]data.StoreInfo, 1)
	branches[0] = data.StoreInfo{Name: "Dominos 1", Latitude: -35.018596, Longitude: 117.883642}
	return branches
}

func (store Dominos) GetName() string {
	return "Dominos"
}
