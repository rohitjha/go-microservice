package main

import (
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	log.Println("Hello, World!")
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
