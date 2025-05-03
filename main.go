package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Xin chào từ Go server!")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server chạy tại http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
