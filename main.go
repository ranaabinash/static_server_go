package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(rw http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(rw, "404 Page not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(rw, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(rw, "Hello Client!!")
}

func formHandler(rw http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(rw, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(rw, "POST requested successfully")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(rw, "Name = %s", name)
	fmt.Fprintf(rw, "Address = %s", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Start of the server")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		log.Fatal(err)
	}
}
