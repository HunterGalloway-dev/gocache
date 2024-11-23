package store

import "gocache/pkg/model"

// TO DO - Implement the KVStore struct
// Defines the basic functions that a store should implement
type PersonStore interface {
	InsertPerson(p model.Person)
	InsertPersons(p []model.Person)
	GetPerson(id int) (model.Person, bool)
	GetAllPersons() []model.Person
	DeletePerson(id int) error
	UpdatePerson(id int, p model.Person) error
	Query(name, email string, ages []int) []model.Person
	String() string
}
