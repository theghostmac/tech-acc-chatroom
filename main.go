package main

import (
	"html/template"
	"net/http"
	// "github.com/gorilla/mux"
)

/*********

- Try to figure out a simple architecture
- Then build on it and improve the quality of the code and arch
- Try to separate the DB concerns and services, handlers etc.
- Utilize interfaces too, it helps a lot

You're doing well ...
**/

func main() {
	// create a multiplexer
	mux := http.NewServeMux()
	// serve static files
	files := http.FileServer(http.Dir("/assets"))
	mux.Handle("/static", http.StripPrefix("/static/", files))
	// redirect the root url to the handler function
	mux.HandleFunc("/", index)
	// other necessary functions to handle
	mux.HandleFunc("/err", err)

	// authentication functions
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/logout", logout)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("signup_account", signupAccount)
	mux.HandleFunc("/authenticate", authenticate)

	// threads and posts functions
	mux.HandleFunc("thread/new", newThread)
	mux.HandleFunc("thread/create", createThread)
	mux.HandleFunc("thread/post", postThread)
	mux.HandleFunc("threa/read", readThread)

	// web server logic
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

// index handler function generates the HTML and write to the ResponseWriter
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
