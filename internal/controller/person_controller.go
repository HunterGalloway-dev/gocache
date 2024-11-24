package controller

import (
	"fmt"
	"gocache/internal/datasource"
	"gocache/pkg/model"
	"gocache/pkg/store"
	"log"
)

// PersonController defines the interface for the person controller
type PersonController interface {
	Health() map[string]string
	GetAllPersons() ([]model.Person, error)
	Query(name, email string, ages []int) ([]model.Person, error)
	UpdatePerson(p model.Person) error
}

// personController is the concrete implementation of PersonController
type personController struct {
	db datasource.DataSource
	kv store.PersonStore // add the data source for the key-value storeh
}

// NewPersonController creates a new instance of personController
func NewPersonController(db datasource.DataSource) (PersonController, error) {
	kv := store.NewKVStore()
	p, err := db.GetAllPersons()
	if err != nil {
		return nil, fmt.Errorf("error getting persons from data source: %w", err)
	}
	kv.InsertPersons(p)
	return &personController{db: db, kv: kv}, nil
}

func (c *personController) Health() map[string]string {
	return c.db.Health()
}

// Query retrieves persons from the data source based on the provided criteria
func (c *personController) Query(name, email string, ages []int) ([]model.Person, error) {
	log.Printf("CONTROLLER: Query called with name=%v, email=%v, ages=%v", name, email, ages)
	p := c.kv.Query(name, email, ages)
	log.Printf("CONTROLLER: Query success: found %v persons", len(p))
	return p, nil
}

// GetAllPersons retrieves all persons from the data source
func (c *personController) GetAllPersons() ([]model.Person, error) {
	log.Printf("CONTROLLER: GetAllPersons called")
	p := c.kv.GetAllPersons()

	log.Printf("CONTROLLER: GetAllPersons success: found %v persons", len(p))

	return p, nil
}

// UpdatePerson updates a person in the data source
func (c *personController) UpdatePerson(p model.Person) error {
	log.Printf("CONTROLLER: UpdatePerson called with person=%v", p)
	err := c.db.UpdatePerson(p)
	if err != nil {
		log.Printf("CONTROLLER: Error updating person: %v", err)
		return err
	}

	// Update the key-value store
	err = c.kv.UpdatePerson(p)
	if err != nil {
		log.Printf("CONTROLLER: Error updating key-value store: %v", err)
		return err
	}

	log.Printf("CONTROLLER: UpdatePerson success")
	return nil
}
