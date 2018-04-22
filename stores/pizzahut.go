package stores

import (
	"../models"
)

type PizzaHut struct {
}

func (store PizzaHut) GetBranchStores() []data.StoreInfo {
	branches := make([]data.StoreInfo, 1)
	branches[0] = data.StoreInfo{Name: "PizzaHut 1"}
	return branches
}

func (store PizzaHut) GetName() string {
	return "PizzaHut"
}
