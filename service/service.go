package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is system service handler")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting Go services on port 8080...")
	http.ListenAndServe(":8080", nil)
}
