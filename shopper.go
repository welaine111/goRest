package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// define shopper
type Shopper struct {
	ID               string //LONG INT
	Location         string
	Firstname        string
	Lastname         string
	Address          *Address
	Email            string
	Cart             *Cart
	RegistrationDate string
}
type Address struct {
	City    string
	State   string
	Address string
}

//print out all items in shoppers array
func GetShoppers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(shoppers)
}

//find shopper from Array shoppers using ID
func GetShopper(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["shopperID"]
	shopper, prs := shoppers[id]
	if prs {
		json.NewEncoder(w).Encode(shopper)
		return
	}

}
