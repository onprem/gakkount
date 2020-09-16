package api

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (a *API) createToken(email, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return at.SignedString([]byte(a.jwtSecret))
}

func (a *API) verifyToken(tok string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tok, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %w", token.Header["alg"])
		}
		return []byte(a.jwtSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
