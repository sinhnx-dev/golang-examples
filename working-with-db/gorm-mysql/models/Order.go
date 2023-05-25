package models

import (
	"time"
)

// Customer type
type Customer struct {
	CustomerID      uint   `json:"customerId" gorm:"column:customer_id;primaryKey;autoIncrement"`
	CustomerName    string `json:"customerName" gorm:"column:customer_name;not null"`
	CustomerAddress string `json:"customerAddress" gorm:"column:customer_address"`
}

// Item type
type Item struct {
	ItemID          uint    `json:"itemId" gorm:"column:item_id;primaryKey;autoIncrement"`
	ItemName        string  `json:"itemName" gorm:"column:item_name"`
	UnitPrice       float64 `json:"unitPrice" gorm:"column:unit_price"`
	Amount          int32   `json:"amount" gorm:"column:amount"`
	ItemStatus      int16   `json:"status" gorm:"column:item_status"`
	ItemDescription string  `json:"description" gorm:"column:item_description"`
	// CreateAt        time.Time `json:"createAt" gorm:"autoCreateTime"`
	// UpdateAt        time.Time `json:"updateAt" gorm:"autoUpdateTime"`
}

// Order type
type Order struct {
	OrderID    uint      `json:"orderId" gorm:"column:order_id;primaryKey;autoIncrement"`
	OrderDate  time.Time `json:"orderDate" gorm:"column:order_date;autoCreateTime"`
	Status     uint      `json:"orderStatus" gorm:"column:order_status"`
	CustomerID uint      `gorm:"column:customer_id"`
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID;references:customer_id"`
	Items      []Item    `json:"items" gorm:"many2many:Orderdetails;foreignKey:OrderID;joinForeignKey:order_id;joinReferences:order_id"`
}
