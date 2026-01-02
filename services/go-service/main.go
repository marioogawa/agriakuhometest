package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Service 🚀")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Go service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
