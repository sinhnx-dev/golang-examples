package model

type Item struct {
	ItemId          int64   `json:"itemId"`
	ItemName        string  `json:"itemName"`
	UnitPrice       float64 `json:"unitPrice"`
	Amount          int32   `json:"amount"`
	ItemStatus      int16   `json:"status"`
	ItemDescription string  `json:"description"`
}
