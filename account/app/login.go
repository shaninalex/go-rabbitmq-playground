package app

import (
	"errors"
	"log"

	"github.com/alexedwards/argon2id"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginPayload) Login(passwodHash string, user_id int64) (string, error) {
	match, err := argon2id.ComparePasswordAndHash(l.Password, passwodHash)
	if err != nil {
		return "", err
	}
	if match {
		return "", errors.New("passwords does ont match")
	}

	log.Println(passwodHash, user_id)
	token, err := CreateJWT(user_id, l.Email)
	if err != nil {
		return "", err
	}

	return token, err
}
