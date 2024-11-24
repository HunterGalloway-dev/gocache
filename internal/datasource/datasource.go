package datasource

import "gocache/pkg/model"

type DataSource interface {
	Health() map[string]string
	GetAllPersons() ([]model.Person, error)
	UpdatePerson(p model.Person) error
}
