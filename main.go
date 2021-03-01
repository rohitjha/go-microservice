package main

import (
	"log"
	"net/http"
)

func handler(rw http.ResponseWriter, r *http.Request) {
	message := "Hello, World!\n"
	log.Print(message)
	rw.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Starting Server")
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}
