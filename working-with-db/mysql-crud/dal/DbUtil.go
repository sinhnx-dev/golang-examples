package dal

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbUtil struct {
	db       *sql.DB
	dbType   string
	hostName string
	port     int
	dbName   string
	userName string
	password string
}

var dbUtil DbUtil

// InitializeMySQL to OrderDB
func InitializeMySQL() {
	dbUtil.dbName = "mysql"
	dbUtil.hostName = "127.0.0.1"
	dbUtil.port = 3306
	dbUtil.dbName = "OrderDB"
	dbUtil.userName = "sinhnx"
	dbUtil.password = "sinhnx.dev"
	dBConnection, err := sql.Open("mysql", "sinhnx:sinhnx.dev@tcp(127.0.0.1:3306)/OrderDB")
	if err != nil {
		fmt.Println("Connection Failed!!")
	}
	err = dBConnection.Ping()
	if err != nil {
		fmt.Println("Ping Failed!!")
	}
	dbUtil.db = dBConnection
	dBConnection.SetMaxOpenConns(10)
	dBConnection.SetMaxIdleConns(5)
	dBConnection.SetConnMaxLifetime(time.Second * 10)
}

// GetConnection is get MySQL Connection
func GetConnection() *sql.DB {
	if dbUtil.db == nil {
		InitializeMySQL()
	}
	return dbUtil.db
}

// CloseStmt after run stmt
func CloseStmt(stmt *sql.Stmt) {
	if stmt != nil {
		stmt.Close()
	}
}

// CloseRows when select
func CloseRows(rows *sql.Rows) {
	if rows != nil {
		rows.Close()
	}
}
