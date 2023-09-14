package db

import "time"

type Order struct {
	Id         int64        `json:"id"`
	UserId     int64        `json:"user_id"`
	OrderCode  string       `json:"order_code"`
	TotalPrice float64      `json:"total_price"`
	CreatedAt  *time.Time   `json:"created_at"`
	Products   OrderProduct `json:"products"`
}

type OrderProduct struct {
	ProductId    int64   `json:"product_id"`
	ProductName  string  `json:"product_name"`
	ProductPrice float64 `json:"product_price"`
	TotalPrice   float64 `json:"total_price"`
	Quantity     int64   `json:"quantity"`
}
