package main

import (
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
