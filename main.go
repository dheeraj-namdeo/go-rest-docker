package main

import (
	"fmt"
	"html"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello,%q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(rw http.ResponseWriter, r *http.Request) { fmt.Fprintf(rw, "Hi, Docker API ") })
	http.ListenAndServe(":8081", nil)
}
