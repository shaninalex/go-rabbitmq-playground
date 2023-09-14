package db

import (
	"database/sql"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func InitializeDatabase(dsn string) (*Database, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	database := &Database{
		DB: db,
	}

	return database, nil
}

// at this moment we know only user_id and product id and quantity
// We need to know products information, calc pricing etc.
func (db *Database) GetOrderInformation(userId int64, products []OrderProduct) (*Order, error) {
	q := goqu.From("users").Select("email", "name").Where(goqu.C("id").Eq(userId))
	selectUser, _, _ := q.ToSQL()
	var user_email, user_name string
	if err := db.DB.QueryRow(selectUser).Scan(
		&user_email, &user_name,
	); err != nil {
		return nil, err
	}

	var order Order
	order.UserId = userId

	return &order, nil
}

func (db *Database) getProductsByIds(products []OrderProduct) error {
	ids := make([]int64, len(products))
	for _, p := range products {
		ids = append(ids, p.ProductId)
	}
	q := goqu.From("products").Select("name", "price").Where(
		goqu.Ex{
			"id": goqu.Op{"in": ids},
		},
	)
	selectProducts, _, _ := q.ToSQL()
	fmt.Println(selectProducts)

	return nil
}
