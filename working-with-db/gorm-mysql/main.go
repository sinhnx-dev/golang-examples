package main

import (
	"fmt"

	"sinhnx.dev/gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// dsn := "sinhnx:sinhnx.dev@tcp(127.0.0.1:3306)/OrderDB?charset=utf8"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// } else {
	// 	fmt.Println("Connect Successfull.")
	// }
	// fmt.Println(db)

	dsn := "sinhnx:sinhnx.dev@tcp(127.0.0.1:3306)/OrderDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connect Successfull.")
	item := models.Item{ItemName: "ITEM 1101", UnitPrice: 12.5, Amount: 10, ItemStatus: 1, ItemDescription: "description..."}
	result := db.Create(&item)
	if result.RowsAffected > 0 {
		fmt.Printf("New item id: %d\n", item.ItemID)
	}

	item.ItemName = "Item 198X"
	result = db.Save(&item)
	if result.RowsAffected > 0 {
		fmt.Printf("update item id: %d, item name: %s\n", item.ItemID, item.ItemName)
	}

	result = db.First(&item, item.ItemID)
	if result.RowsAffected > 0 {
		fmt.Println(item)
	}

	result = db.Delete(&models.Item{}, item.ItemID)
	if result.RowsAffected > 0 {
		fmt.Printf("deleted item id: %d\n", item.ItemID)
	}

	//create order
	customer := models.Customer{CustomerName: "Nguyen A", CustomerAddress: "Hanoi"}
	if db.Create(&customer).RowsAffected > 1 {
		fmt.Printf("Inserted customer id: %d\n", customer.CustomerID)

		db.Begin()
		items := []models.Item{
			models.Item{ItemID: 1, Amount: 2},
			models.Item{ItemID: 2, Amount: 1},
		}

		order := models.Order{Customer: customer, Status: 1, Items: items}
		if db.Create(&order).RowsAffected < 1 {
			db.Rollback()
		}
		db.Commit()

		fmt.Printf("Inserted order id: %d\n", order.OrderID)
	}
}
