package server

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/aagueromeli/movies_api/internal/application"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

type ServerChi struct {
	ServerAddr string
}

func LoadServerConf() (*ServerChi, error) {
	err := godotenv.Load()
	if err != nil {
		msg := fmt.Sprintf("Failed to load .env: %s", err)
		return nil, errors.New(msg)
	}
	// default values
	serverAddr := os.Getenv("ADDRESS")

	if serverAddr == "" {
		serverAddr = ":8080"
	}

	return &ServerChi{
		ServerAddr: serverAddr,
	}, nil
}

func (c *ServerChi) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	healthcheck := application.HealthRouter()

	router.Mount("/healthcheck", healthcheck)

	err := http.ListenAndServe(c.ServerAddr, router)
	return err
}
