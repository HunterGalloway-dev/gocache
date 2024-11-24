package datasource

import (
	"gocache/pkg/model"
)

type MockDataSource struct {
	persons []model.Person
}

// NewMockDataSource creates a new instance of MockDataSource this is a local in-memory data source for testing
func NewMockDataSource() DataSource {
	persons := []model.Person{
		{ID: 1, Name: "John Doe", Age: 30, Email: "john.doe@example.com"},
		{ID: 2, Name: "Jane Smith", Age: 25, Email: "jane.smith@example.com"},
	}
	return &MockDataSource{
		persons: persons,
	}
}

func (m *MockDataSource) Health() map[string]string {
	return map[string]string{"status": "healthy"}
}

func (m *MockDataSource) GetAllPersons() ([]model.Person, error) {
	return m.persons, nil
}

func (m *MockDataSource) UpdatePerson(p model.Person) error {
	for i, person := range m.persons {
		if person.ID == p.ID {
			m.persons[i] = p
			return nil
		}
	}
	return nil
}
