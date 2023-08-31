package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	TOKEN_TIME   = 5 * time.Minute
	REFRESH_TIME = 1 * time.Hour
	SECRET       = "secret"
)

type User struct {
	Name   string
	Age    int
	Email  string
	UserId string
}

type UserClaims struct {
	UserId string
	jwt.StandardClaims
}

func main() {

}

func GenerateTokens(user *User) (string, string, error) {
	var err error

	tokenTime := time.Now().Add(TOKEN_TIME).Unix()
	refreshTime := time.Now().Add(REFRESH_TIME).Unix()

	userClaims := getClaimsFromUser(user)
	userClaims.StandardClaims = jwt.StandardClaims{
		ExpiresAt: tokenTime,
	}
	refreshClaims := jwt.StandardClaims{
		ExpiresAt: refreshTime,
	}

	Token, _ := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString([]byte(SECRET))
	RefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET))
	if err != nil {
		log.Println("error while signing token : " + err.Error())
		return "", "", err
	}
	return Token, RefreshToken, nil
}

func GetClaimsFromToken(token string, claims jwt.Claims) error {
	token_parsed, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})

	if err != nil || !token_parsed.Valid {
		log.Println("unauthorized : " + err.Error())
		return err
	}
	return nil
}

func getClaimsFromUser(user *User) *UserClaims {
	uc := new(UserClaims)
	uc.UserId = user.UserId
	return uc
}
