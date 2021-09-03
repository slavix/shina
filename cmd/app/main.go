package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world\n")
	})

	fmt.Println("Server listening!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
