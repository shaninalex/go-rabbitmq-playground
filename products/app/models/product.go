package models

import (
	"math/rand"
	"strings"
	"time"
)

type Product struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	ProductCode string    `json:"product_code"`
	Price       float64   `json:"price" binding:"required"`
	CreatedAt   time.Time `json:"created_at"`
}

func (p *Product) Init() {
	p.ProductCode = randomString(8)
	p.CreatedAt = time.Now()
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
