package main

import (
	"fmt"
	"log"
	"net/http"
	"shina/internal/site"
	"shina/internal/utils"
	"shina/pkg/config"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderHTML(w, r, "home", &site.HTMLData{})
	})

	fmt.Println("Server listening!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
