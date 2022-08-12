package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	// user, _ := data.U
}

type DataToEncode struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

func AuthUser() {

	// let's say you've authenticated the user
	// we want to generate cookies, but we need a token
	// preferably jwt tokens

	expiration_date := time.Now().Add(5 * time.Minute)

}
