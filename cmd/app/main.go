package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/ztrue/tracerr"
	"os"
	"os/signal"
	"shina/internal/handlers"
	"shina/internal/repository"
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

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	fmt.Println(db)

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
