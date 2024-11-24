package server

import (
	"gocache/pkg/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) healthHandler(c *gin.Context) {
	dbHealth := s.pc.Health()

	health := make(map[string]interface{})
	health["app"] = "OK"
	health["database"] = dbHealth

	c.JSON(http.StatusOK, health)
}

func (s *Server) getPersonsHandler(c *gin.Context) {
	log.Printf("ROUTE: getPersonsHandler called: %v %v ", c.Request.Method, c.Request.URL.Path)
	persons, err := s.pc.GetAllPersons()
	if err != nil {
		log.Printf("ROUTE: getPersonsHandler error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ROUTE: getPersonsHandler success: found %v persons", len(persons))

	c.JSON(http.StatusOK, persons)
}

func (s *Server) queryPersonsHandler(c *gin.Context) {

	name := c.Query("name")
	email := c.Query("email")
	ageStr := c.QueryArray("ages")
	log.Printf("ROUTE: queryPersonsHandler called: %v %v name=%v, email=%v, ages=%v", c.Request.Method, c.Request.URL.Path, name, email, ageStr)
	ages, err := stringSliceToIntSlice(ageStr)

	if err != nil {
		log.Printf("ROUTE: queryPersonsHandler error converting string slice to int slice: %v", err)
	}

	log.Printf("ROUTE: queryPersonsHandler called: %v %v name=%v, email=%v, ages=%v", c.Request.Method, c.Request.URL.Path, name, email, ages)

	persons, err := s.pc.Query(name, email, ages)
	if err != nil {
		log.Printf("ROUTE: queryPersonsHandler error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ROUTE: queryPersonsHandler success: found %v persons", len(persons))

	c.JSON(http.StatusOK, persons)
}

func (s *Server) updatePersonHandler(c *gin.Context) {
	log.Printf("ROUTE: updatePersonHandler called: %v %v", c.Request.Method, c.Request.URL.Path)
	var person model.Person
	if err := c.BindJSON(&person); err != nil {
		log.Printf("ROUTE: updatePersonHandler error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := s.pc.UpdatePerson(person)
	if err != nil {
		log.Printf("ROUTE: updatePersonHandler error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("ROUTE: updatePersonHandler success")

	c.JSON(http.StatusOK, gin.H{"message": "Person updated successfully"})
}

// Given a slice of strings, convert them to a slice of integers, if conversion fails return an error
func stringSliceToIntSlice(strSlice []string) ([]int, error) {
	log.Printf("ROUTE: Converting string slice to int slice: %v", strSlice)
	intSlice := make([]int, 0, len(strSlice))

	if len(strSlice) == 0 {
		return intSlice, nil
	}

	if len(strSlice) == 1 && strSlice[0] == "" {
		return intSlice, nil
	}

	for _, str := range strSlice {
		intVal, err := strconv.Atoi(str)
		if err != nil {
			return intSlice, err
		}
		intSlice = append(intSlice, intVal)
	}
	return intSlice, nil
}