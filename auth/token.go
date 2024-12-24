package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/humblgod/belajar-golang-rest-api/config"
)

 func CreateJWT(secret []byte, UserId int) (string, error) {
	// expiration time define 
	expiration := time.Second * time.Duration(config.Envs.JWTExpirationTime)

	// create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"UserId" : strconv.Itoa(UserId),
		"expiredAt" : time.Now().Add(expiration).Unix(),
	})

	// sign the token and return token as string ""
	stringToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return stringToken, nil
 }