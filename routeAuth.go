package main

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	// user, _ := data.U
}

// to create a token you will need some data to be
// encoded inside the generated jwt token
// so that you will be able to call the fields from the claims

type DataToEncode struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

func AuthUser() {

	var (
		user_data = User{}
	)

	// let's say you've authenticated the user
	// we want to generate cookies, but we need a token
	// preferably jwt tokens

	expiration_date := time.Now().Add(5 * time.Minute)

	encode_data := DataToEncode{
		Username: user_data.Username,
	}

}
