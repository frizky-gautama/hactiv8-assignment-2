package model

import "time"

type Orders struct {
	OrderID      int       `json:"order_id"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
	Item         []Items   `json:"items"`
}

type Items struct {
	ItemID      int    `json:"item_id" example:"1"`
	ItemCode    string `json:"item_code" example:"ITEM01"`
	Description string `json:"description" example:"DESC ITEM 01"`
	Quantity    int    `json:"quantity" example:"10"`
	OrderId     int    `json:"order_id" example:"1"`
}

type Request struct {
	CustomerName string  `json:"customer_name" example:"Ricky"`
	Item         []Items `json:"items"`
}

type Response struct {
	CustomerName string    `json:"customer_name" example:"Ricky"`
	OrderID      int       `json:"order_id" example:"1"`
	OrderedAt    time.Time `json:"ordered_at" example:"2022-10-04T11:07:59.6868076+07:00"`
	Item         []Items   `json:"items"`
}
