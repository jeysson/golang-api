package main

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jeysson/golang-api/models"
)

var jwtkey = []byte("your_secret_key")

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &models.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidationToken(tokenString string) (*models.Claims, error) {
	claims := &models.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, err
		}

		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
