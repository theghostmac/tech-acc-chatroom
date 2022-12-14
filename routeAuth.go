package main

import (
	"errors"
	"log"
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

func AuthUser(w http.ResponseWriter, r *http.Request) {

	var (
		user_data             = User{}
		jwt_secret_key        = []byte("secret_key")
		ErrCouldNotSignString = "could not sign jwt token"
	)

	// let's say you've authenticated the user
	// we want to generate cookies, but we need a token
	// preferably jwt tokens

	expiration_date := time.Now().Add(5 * time.Minute)

	// remember we need to encode some data
	encode_data := DataToEncode{
		Username: user_data.Username,
		Email:    user_data.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiration_date.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "tech-acc-chatroom",
		},
	}

	// signing method is HS256 because if you choose RS256
	// you will need a private key generated by openssl to be passed in later

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, encode_data)
	token_string, err := token.SignedString(jwt_secret_key)
	if err != nil {
		log.Print(errors.New(ErrCouldNotSignString))
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "user-token",
		Value:   token_string,
		Expires: expiration_date,
	})

	// now the cookie has been generated, you might want to write a middleware that
	// checks when the cookie is about to expire and refresh it if the user is still active.
}
