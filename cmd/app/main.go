package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/ztrue/tracerr"
	"os"
	"os/signal"
	"shina/internal/handlers"
	"shina/internal/server"
	"shina/pkg/config"
	"shina/pkg/logger"
	"syscall"
)

func main() {
	logger.Init("shina-web-server", 5)

	if err := config.InitConfig(); err != nil {
		logger.Panic(err, "error initializing configs")
	}

	if err := godotenv.Load(); err != nil {
		logger.Panic(err, "error loading env variables")
	}

	handler := handlers.NewHandler()

	srv := new(server.Server)
	go func() {
		if err := srv.Run(config.GetString("site.port"), handler.InitRoutes()); err != nil {
			logger.Panic(err, "error occured while running http server")
		}
	}()

	fmt.Println("Server listening at port " + config.GetString("site.port"))

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := shutdown(srv); err != nil {
		logger.Panic(err, "shutdown failed")
	}
}

func shutdown(server *server.Server) error {
	if err := server.Shutdown(context.Background()); err != nil {
		return tracerr.Wrap(errors.New("http server doesn't close connection"))
	}

	return nil
}
