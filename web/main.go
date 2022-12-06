package main

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Go world !</h1>")
	fmt.Fprintf(w, "<p>Golang is a wonderful language</p>")
	// http.ServeFile(w, r, "./auth/signup.html")
	// http.Handle("/", http.FileServer(http.Dir("./front-end")))

}

func signup_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Go world !</h1>")
	// fmt.Fprintf(w, "<p>Golang is a wonderful language</p>")
	http.ServeFile(w, r, "./auth/signup.html")

}

func login_handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "<h1>Go world !</h1>")
	// fmt.Fprintf(w, "<p>Golang is a wonderful language</p>")
	http.ServeFile(w, r, "./auth/login.html")

}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<h1>About Golang</h1>
	<p>Golang is an open source programming language which allows you to build simple, reliable, and efficient applications.</p>
	`)
	http.FileServer(http.Dir("./front-end"))
}

func main() {
	http.HandleFunc("/", index_handler)
	//http.handle
	// http.Handle("/", http.FileServer(http.Dir("./front-end")))
	http.HandleFunc("/register/", signup_handler)
	http.HandleFunc("/login/", login_handler)
	http.ListenAndServe(":9786", nil)
}
