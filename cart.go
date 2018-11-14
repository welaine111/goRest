package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CartItem struct {
	CartItemID string
	Coupon     string
	Amount     int
	Product    *Product
}

type Product struct {
	ProductID   int64
	Name        string
	Image       string
	Price       float32
	Category    string
	Weight      float32
	Size        string
	Description string
	shelfIDList []string
}

type SendProduct struct {
	ProductID   string
	Name        string
	Description string
	Coupon      string
	CartItemID  string
	amount      int
}

//type CartList map[int64]*CartItem

type Cart struct {
	ID        string
	CartItems map[string]*CartItem
}

//func NewCart() *Cart{
//	var c Cart

//}

func getCart(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shopperId := params["shopperID"]
	cart := shoppers[shopperId].Cart
	json.NewEncoder(w).Encode(cart.CartItems)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shopperId := params["shopperID"]
	itemId := params["itemID"]

	cart := shoppers[shopperId].Cart
	item, prs := cart.CartItems[itemId]
	if prs {

		json.NewEncoder(w).Encode(item)
		return
	}

}
