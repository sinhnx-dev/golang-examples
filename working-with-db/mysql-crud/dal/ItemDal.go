package dal

import (
	"github.com/sinhnx-dev/golang-examples/working-with-db/mysql-crud/model"
)

// InsertItem to OrderDB
func InsertItem(item model.Item) (int64, int64, error) {
	InitializeMySQL()
	db := GetConnection()
	sqlQuery := "INSERT Items SET item_name=?, unit_price=?, amount=?, item_status=?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, 0, err
	}
	res, err := stmt.Exec(item.ItemName, item.UnitPrice, item.Amount, item.ItemStatus)
	if err != nil {
		return 0, 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return rowsAffected, lastInsertId, err
}

// UpdateItem in OrderDB
func UpdateItem(item model.Item) (int64, error) {
	db := GetConnection()
	sqlQuery := "UPDATE Items SET item_name=?, unit_price=?, amount=?, item_status=? WHERE item_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(item.ItemName, item.UnitPrice, item.Amount, item.ItemStatus, item.ItemId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

// DeleteItem in OrderDB with parameter is itemId
func DeleteItem(itemId int64) (int64, error) {
	db := GetConnection()
	defer db.Close()
	sqlQuery := "DELETE FROM Items WHERE item_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(itemId)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, err
}

// GetItem from itemId
func GetItem(itemId int64) (model.Item, error) {
	db := GetConnection()
	sqlQuery := "SELECT item_id, item_name, unit_price, amount, item_status, item_description FROM Items WHERE item_id = ?"
	stmt, err := db.Prepare(sqlQuery)
	defer CloseStmt(stmt)
	var item model.Item
	if err != nil {
		return item, err
	}
	res, err := stmt.Query(itemId)
	defer CloseRows(res)
	if err != nil {
		return item, err
	}
	if res.Next() {
		res.Scan(&item.ItemId, &item.ItemName, &item.UnitPrice, &item.Amount, &item.ItemStatus, &item.ItemDescription)
	}
	return item, err
}

// func InsertTransaction() {
// 	db := GetConnection()
// 	tx, err := db.Begin()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer tx.Rollback()
// 	stmt, err := tx.Prepare("INSERT INTO Items VALUES (?, ?, ?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	_, err := stmt.Exec(itemName, unitPrice, amount, itemStatus)
// 	// more stmt.Exec() ...

// 	err = tx.Commit()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stmt.Close()

// 	for row.Next() {
// 		// ...
// 	}
// 	if err = row.Err(); err != nil {
// 		// handle the error here
// 	}

// 	row, err := db.QueryRow("select item_id, item_name, unit_price from items where item_id = ?", 5)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			// there were no rows, but otherwise no error occurred
// 		} else {
// 			log.Fatal(err)
// 		}
// 	}

// 	if driverErr, ok := err.(*mysql.MySQLError); ok {
// 		// Now the error number is accessbile directly
// 		if driverErr.Number == 1045 {
// 			// Handle the permission denied error
// 		}
// 	}

// 	if driverErr, ok := err.(*mysql.MySQLError); ok {
// 		if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
// 			// Handle the permission denied error
// 		}
// 	}

// 	for rows.Next() {
// 		var s sql.NullString
// 		err := rows.Scan(&s)
// 		if s.Valid {
// 			// s is a string
// 		}else {
// 			// s is NULL
// 		}
// 	}

// 	rows, err := db.Query("SELECT COALESCE(item_description, '') description FROM Items WHERE item_id=?")
// 	for rows.Next() {
// 		err := rows.Scan(&description)
// 		// If `item_description` was NULL, `description` is an empty string
// 		// This works with other data types as well
// 	}

// 	cols, err := rows.Columns()
// 	if err != nil {
// 		// handling error
// 	} else {
// 		dest := [] interface {} { // standard MySQL columns
// 			new(uint64), 	// id
// 			new(string),	// host
// 			new(string), 	// user
// 			new(string), 	// db
// 			new(string), 	// command
// 			new(uint32), 	// time
// 			new(string)		// state
// 			new(string), 	// info
// 		}
// 		if len(cols) == 11 {
// 			// Percona Server
// 		} else if len(cols) > 8 {
// 			// handle this case
// 		}
// 		err = rows.Scan(dest...)
// 		// Work with the values in dest
// 	}

// 	cols, err := rows.Columns() // Remember to check err afterwards
// 	vals := make([]interface{}, len(cols))
// 	for i, _ := range cols {
// 		vals[i] = new(sql.RawBytes)
// 	}
// 	for rows.Next() {
// 		err = rows.Scan(vals...)
// 		// Now you can check each element of vals for nil-ness and
// 		// you can use type introspection and type assertions to fetch the column into typed variable.
// 	}
// }
