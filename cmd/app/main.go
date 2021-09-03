package main

import (
	"fmt"
	"log"
	"net/http"
	"shina/pkg/config"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello world\n")
	})

	fmt.Println("Server listening!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
