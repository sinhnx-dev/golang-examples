package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	ItemId    int     `json:"itemId"`
	ItemName  string  `json:"name"`
	UnitPrice float32 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
	Desc      string  `json:"desc"`
}

var Items []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	// add our items route and map it to our returnAllItems function like so
	http.HandleFunc("/items", returnAllItems)
	log.Fatal(http.ListenAndServe(":1411", nil))
}

func main() {
	Items = []Item{
		{ItemId: 1, ItemName: "Item 1", UnitPrice: 12.5, Quantity: 5, Desc: "Item 1 Description"},
		{ItemId: 2, ItemName: "Item 2", UnitPrice: 15.5, Quantity: 3, Desc: "Item 2 Description"},
		{ItemId: 3, ItemName: "Item 3", UnitPrice: 18.5, Quantity: 2, Desc: "Item 3 Description"},
	}
	handleRequests()
}
func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /items -> Endpoint Hit: returnAllItems")
	json.NewEncoder(w).Encode(Items)
}
