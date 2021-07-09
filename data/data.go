package data

import (
	"vendingMachine/models"
	"vendingMachine/roles"
)

var UserList = make([]models.UserModel, 0)
var AcceptedCoins = [5]int{5, 10, 20, 50, 100}

var AvailableProducts = make([]models.ProductModel, 0)

var BuyersList []roles.Buyer
var SellersList []roles.Seller

func getNextID(i interface{}) int {
	var v int
	switch i.(type) {
	case models.UserModel:
		last := UserList[len(UserList)-1]
		v = last.ID + 1
	case models.ProductModel:
		last := AvailableProducts[len(AvailableProducts)-1]
		v = last.ID + 1
	default:
		v = -1
	}
	return v
}
