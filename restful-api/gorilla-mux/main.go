package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct {
	ItemId    int     `json:"itemId"`
	ItemName  string  `json:"itemName"`
	UnitPrice float32 `json:"unitPrice"`
	Quantity  int     `json:"quantity"`
	Desc      string  `json:"desc"`
}

var Items []Item

func main() {
	fmt.Println("Rest API - Mux Routers")
	Items = []Item{
		{ItemId: 1, ItemName: "Item 1", UnitPrice: 12.5, Quantity: 5, Desc: "Item 1 Description"},
		{ItemId: 2, ItemName: "Item 2", UnitPrice: 15.5, Quantity: 3, Desc: "Item 2 Description"},
		{ItemId: 3, ItemName: "Item 3", UnitPrice: 18.5, Quantity: 2, Desc: "Item 3 Description"},
	}
	handleRequests()
}
func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/item", createNewItem).Methods("POST")
	myRouter.HandleFunc("/items", returnAllItems)
	myRouter.HandleFunc("/item/{id}", returnSingleItem)
	myRouter.HandleFunc("/item/update/{id}", updateItem).Methods("PUT")
	myRouter.HandleFunc("/item/delete/{id}", deleteItem).Methods("DELETE")
	// finally, instead of passing in nil, we want to pass in our newly created router as the second argument
	log.Fatal(http.ListenAndServe(":2111", myRouter))
}
func returnAllItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET: /items -> Endpoint Hit: returnAllItems")
	json.NewEncoder(w).Encode(Items)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func returnSingleItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("READ ITEM...")
	vars := mux.Vars(r)
	key := vars["id"]
	id, err := strconv.Atoi(key)
	if err != nil {
		fmt.Println(err)
	}

	for _, item := range Items {
		if item.ItemId == id {
			json.NewEncoder(w).Encode(item)
		}
	}
}
func createNewItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CREATE ITEM...")
	// get the body of our POST request return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w, "%+v", string(reqBody))
	var item Item
	json.Unmarshal(reqBody, &item)
	// update our global Items array to include our new Item
	Items = append(Items, item)
	json.NewEncoder(w).Encode(item)
}
func updateItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UPDATE ITEM...")
	// get the body of our POST request return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to delete
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	var item Item
	json.Unmarshal(reqBody, &item)
	// update our global Items array to include our Item
	for index, i := range Items {
		if i.ItemId == id {
			item.ItemId = id
			Items[index] = item
		}
	}
	json.NewEncoder(w).Encode(item)
}
func deleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DELETE ITEM...")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we wish to delete
	idString := vars["id"]
	id, err := strconv.Atoi(idString)
	if err != nil {
		fmt.Println(err)
	}
	// we then need to loop through all our items
	for index, item := range Items {
		// if our id path parameter matches one of our items
		if item.ItemId == id {
			// updates our Items array to remove the item
			Items = append(Items[:index], Items[index+1:]...)

			fmt.Println("delete...")
		}
	}
}
