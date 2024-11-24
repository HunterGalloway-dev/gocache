package server

import (
	"fmt"
	"gocache/internal/controller"
	"gocache/internal/datasource"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	pc controller.PersonController
}

func NewServer() (*http.Server, error) {
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {

		return nil, fmt.Errorf("error converting PORT to integer: %v", err)
	}

	db, err := datasource.NewMongo()

	if err != nil {
		return nil, fmt.Errorf("error creating mongo data source: %v", err)
	}

	// Create controllers
	pc, err := controller.NewPersonController(db)
	if err != nil {
		return nil, fmt.Errorf("error creating person controller: %v", err)
	}

	NewServer := &Server{
		port: port,
		pc:   pc,
	}

	server := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server, nil
}
