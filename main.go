package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start of the server")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server started at localhost:8080")
}
