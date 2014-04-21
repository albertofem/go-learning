package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Other route is here: <pre>%s</pre>", r)
}

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/other-route", handler)
    http.ListenAndServe(":8080", nil)
}