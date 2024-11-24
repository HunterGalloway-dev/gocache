package server

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	r.GET("/health", s.healthHandler)
	r.GET("/persons", s.getPersonsHandler)
	r.GET("/persons/filter", s.queryPersonsHandler)

	r.POST("/persons/update", s.updatePersonHandler)

	return r
}
