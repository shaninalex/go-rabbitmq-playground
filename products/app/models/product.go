package models

import (
	"math/rand"
	"strings"
	"time"
)

//	   id SERIAL PRIMARY KEY,
//     name TEXT NOT NULL UNIQUE,
//     product_code VARCHAR(8) NOT NULL,
//     price FLOAT NOT NULL,
//     quantity INTEGER NOT NULL,
//     created_at TIMESTAMP DEFAULT current_timestamp

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	ProductCode string    `json:"product_code"`
	Price       float64   `json:"price" binding:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Product) Update(name string, price float64) {
	p.Name = name
	p.Price = price
}

func Create(name string, price float64) Product {
	p := Product{}
	p.Name = name
	p.Price = price
	p.ProductCode = randomString(8)
	p.CreatedAt = time.Now()

	return p
}


const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
