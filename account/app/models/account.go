package models

import (
	"account/app/utils"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
)

type NewUser struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (a *NewUser) Create(db *sql.DB) error {
	passwordHash, err := utils.GenerateFromPassword(a.Password)
	if err != nil {
		return err
	}
	ds := goqu.Insert("users").Rows(
		goqu.Record{
			"name":     a.Name,
			"email":    a.Email,
			"password": passwordHash,
		},
	).Returning("id")
	insertSQL, _, _ := ds.ToSQL()
	err = db.QueryRow(insertSQL).Scan(&a.Id)
	if err != nil {
		return err
	}
	return nil
}

type User struct {
	Id       int64  `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"-"`
}

func (a *User) Get(db *sql.DB, email string, user_id int64) error {
	ds := goqu.From("users").Select(
		"id", "name", "email", "password",
	)
	if email != "" {
		ds = ds.Where(goqu.C("email").Eq(email))
	} else {
		ds = ds.Where(goqu.C("id").Eq(user_id))
	}

	insertSQL, _, _ := ds.ToSQL()
	err := db.QueryRow(insertSQL).Scan(
		&a.Id, &a.Name, &a.Email, &a.Password,
	)
	if err != nil {
		return err
	}
	return nil
}
