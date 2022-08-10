package main

import (
	"html/template"
	"net/http"
	// "github.com/gorilla/mux"
)

func main() {
	// create a multiplexer
	mux := http.NewServeMux()
	// serve static files
	files := http.FileServer(http.Dir("/assets"))
	mux.Handle("/static", http.StripPrefix("/static/", files))
	// redirect the root url to the handler function
	mux.HandleFunc("/", index)
	// web server logic
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(writer, "layout", threads)
	}
}
