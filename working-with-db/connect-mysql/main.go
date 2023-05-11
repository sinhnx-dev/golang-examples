package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("Connect to MySQL...")

	// Open up our database connection.
	// The database is called OrderDB
	db, err := sql.Open("mysql", "sinhnx:sinhnx.dev@tcp(127.0.0.1:3306)/OrderDB")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Completed!")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	} else {
		fmt.Println("Successful database connection")
	}

	var (
		id    int
		name  string
		price float64
	)
	rows, err := db.Query("select item_id, item_name, unit_price from items where item_id = ?", 5)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name, &price)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("| %-2d | %-20s | %5.2f |\n", id, name, price)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
