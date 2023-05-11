package main

import (
	"fmt"

	"github.com/sinhnx-dev/golang-examples/working-with-db/mysql-crud/dal"
	"github.com/sinhnx-dev/golang-examples/working-with-db/mysql-crud/model"
)

func main() {
	item := model.Item{0, "Item 100", 15.5, 12, 1, ""}
	rowsAffected, lastInsertedId, err := dal.InsertItem(item)

	if err != nil {
		fmt.Println("Insert error.")
	} else {
		if rowsAffected > 0 {
			fmt.Println("Insert completed.")
			item.ItemId = lastInsertedId
			fmt.Printf("new item id is %d\n", item.ItemId)
		}
	}

	//update item
	item.ItemName = "Item 10001"
	rowsUpdatedAffected, err := dal.UpdateItem(item)
	if err != nil {
		fmt.Println("Update error.")
	} else {
		if rowsUpdatedAffected > 0 {
			fmt.Println("update completed.")
		}
	}

	//get item
	lastInsertedItem, err := dal.GetItem(4)
	if err != nil {
		fmt.Println("get error.")
	} else {
		fmt.Printf("%d - %s - %f - %d\n", lastInsertedItem.ItemId, lastInsertedItem.ItemName, lastInsertedItem.UnitPrice, item.Amount)
	}

	//delete item
	rowsDeletedAffected, err := dal.DeleteItem(lastInsertedId)
	if err != nil {
		fmt.Println("delete error.")
	} else {
		if rowsDeletedAffected > 0 {
			fmt.Println("delete completed.")
		}
	}
}
