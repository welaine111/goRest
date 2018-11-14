package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

//func GetCart()
// our main function
func testJson(code *Shopper) {
	b, err := json.Marshal(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

func testCartItem(code CartItem) {
	b, err := json.Marshal(code)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
}

type Shoppers map[string]*Shopper

var shoppers Shoppers

func main() {
	router := mux.NewRouter()

	shoppers = make(Shoppers)

	//g := &Graph{connections: make(map[Vertex][]Vertex),}
	shoppers["1000"] = &Shopper{ID: "1000", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}, Email: "ww@gmail.com", Cart: &Cart{ID: "11233", CartItems: make(map[string]*CartItem)}, RegistrationDate: "12-01-2018"}
	shoppers["2000"] = &Shopper{ID: "2000", Firstname: "Sara", Lastname: "Vil", Address: &Address{City: "City Z", State: "State Q"}, Email: "123@gmail.com", Cart: &Cart{ID: "11000", CartItems: make(map[string]*CartItem)}, RegistrationDate: "11-01-2015"}
	product1 := Product{ProductID: 9999, Name: "Pepsi", Price: 1.99, Weight: 1.2, Size: "Small", Description: "soda pepsi is yummy!", Category: "Soda"}

	cart1 := CartItem{CartItemID: "1112", Coupon: "ook", Amount: 10, Product: &product1}
	testCartItem(cart1)
	fmt.Println(cart1)

	shoppers["2000"].Cart.CartItems["9999"] = &cart1
	shoppers["2000"].Cart.CartItems["8888"] = &CartItem{Coupon: "buy2 get 1 free", Amount: 15, Product: &product1, CartItemID: "11234"}
	shoppers["2000"].Cart.CartItems["9999"] = &CartItem{Coupon: "buy2 get 1 free", Amount: 2} //å, product: &product1}
	shoppers["2000"].Cart.CartItems["10000"] = new(CartItem)
	fmt.Println(shoppers["2000"].Cart.CartItems["8888"].Coupon)
	fmt.Println(shoppers["2000"].Cart.CartItems["8888"].Product)

	xt := reflect.TypeOf(shoppers["2000"].Cart.CartItems).Kind()
	yt := reflect.TypeOf(shoppers["2000"].Cart.CartItems["8888"].Product).Kind()
	zt := reflect.TypeOf(shoppers["2000"].Cart.CartItems["8888"]).Kind()

	fmt.Printf("%T: %s\n", xt, xt)
	fmt.Printf("%T: %s\n", yt, yt)
	fmt.Printf("%T: %s\n", zt, zt)

	router.HandleFunc("/shoppers/", GetShoppers).Methods("GET")
	router.HandleFunc("/shopper/{shopperID}", GetShopper).Methods("GET")
	router.HandleFunc("/{shopperID}/cart/{cartID}", getCart).Methods("GET")
	router.HandleFunc("/{shopperID}/{cartID}/item/{itemID}", getItem).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))

}
