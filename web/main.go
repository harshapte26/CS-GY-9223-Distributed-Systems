package main

import ("fmt"
	"net/http")

func index_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Go world !</h1>")
	fmt.Fprintf(w, "<p>Golang is a wonderful language</p>")
}

func about_handler(w http.ResponseWriter, r* http.Request){
	fmt.Fprintf(w, `
	<h1>About Golang</h1>
	<p>Golang is an open source programming language which allows you to build simple, reliable, and efficient applications.</p>
	`)
}

func main(){
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.ListenAndServe(":8000", nil)
}