package server

import (
	"gocache/internal/controller"
	"gocache/internal/datasource"
	"log"
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

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	db := datasource.NewMongo()

	// Create controllers
	pc, err := controller.NewPersonController(db)
	if err != nil {
		log.Fatalf("Error creating person controller: %v", err)
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

	return server
}
