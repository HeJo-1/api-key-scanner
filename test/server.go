package main

import (
	"net/http"
	"fmt"
)

func main() {
	fmt.Println("Local test server running on http://localhost:8080")
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}

