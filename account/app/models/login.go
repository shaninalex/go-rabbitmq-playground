package models

import (
	"errors"

	"account/app/utils"
)

type LoginPayload struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (l *LoginPayload) Login(passwodHash string, user_id int64) (string, error) {
	match, err := utils.ComparePasswordAndHash(l.Password, passwodHash)
	if err != nil {
		return "", err
	}
	if !match {
		return "", errors.New("password does not match")
	}

	token, err := utils.CreateJWT(user_id, l.Email)
	if err != nil {
		return "", err
	}

	return token, err
}
