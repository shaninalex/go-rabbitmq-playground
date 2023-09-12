package db

import "time"

type Order struct {
	Id         int64      `json:"id"`
	UserId     int64      `json:"user_id"`
	OrderCode  string     `json:"order_code"`
	TotalPrice string     `json:"total_price"`
	CreatedAt  *time.Time `json:"created_at"`
}

type OrderProduct struct {
	ProductName  string `json:"product_name"`
	ProductPrice string `json:"product_price"`
	TotalPrice   string `json:"total_price"`
	Quantity     string `json:"quantity"`
}
