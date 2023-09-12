package models

import (
	"math/rand"
	"strings"
	"time"
)

type Product struct {
	Id          int64     `json:"id" db:"id"`
	Name        string    `json:"name" binding:"required" db:"name"`
	ProductCode string    `json:"product_code" db:"product_code"`
	Price       float64   `json:"price" binding:"required" db:"price"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

func (p *Product) Init() {
	p.ProductCode = randomString(8)
	p.CreatedAt = time.Now()
}

const charset = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	sb := strings.Builder{}
	sb.Grow(n)
	for i := 0; i < n; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}
