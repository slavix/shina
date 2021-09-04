package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"shina/internal/site"
	"shina/internal/utils"
	"shina/pkg/config"
	"shina/pkg/logger"
)

func main() {
	logger.Init("shina-web-server", 5)

	if err := config.InitConfig(); err != nil {
		logger.Panic(err, "error initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		logger.Panic(err, "error loading env variables")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RenderHTML(w, r, "home", &site.HTMLData{})
	})

	fmt.Println("Server listening!")

	log.Fatal(http.ListenAndServe(":3000", nil))
}
