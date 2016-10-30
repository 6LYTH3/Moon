package main

import (
	"./app/route"
	"log"
	"net/http"
)

func main() {
	log.Printf("Come on baby, I listening on port 8080!")
	// http.Handle("/", http.FileServer(http.Dir("public")))

	log.Fatal(http.ListenAndServe(":8080", route.Load()))
}
