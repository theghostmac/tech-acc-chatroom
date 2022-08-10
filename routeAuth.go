package main

import (
	"net/http"

	"github.com/theghostmac/tech-acc-chatroom/data"
)

func authenticate(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	user, _ :=
}