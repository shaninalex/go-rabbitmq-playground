package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateJWT(id int64, email string) (string, error) {
	// https://www.golinuxcloud.com/golang-jwt/
	log.Println(email, id)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(t string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		log.Println(err)
		log.Println("error on parsing")
		return nil, err
	}

	if token == nil {
		log.Println(err)
		log.Println("invalid token")
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("couldn't parse claims")
		return nil, err
	}

	exp := claims["exp"].(float64)
	if int64(exp) < time.Now().Local().Unix() {
		log.Println("token expired")
		return nil, err
	}

	return claims, nil
}
