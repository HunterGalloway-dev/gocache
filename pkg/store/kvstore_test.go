package store

import (
	"gocache/pkg/model"
	"testing"
)

func TestInsertPerson(t *testing.T) {
	// Setup the KVStore
	store := NewKVStore()

	// Create a test person
	person := model.Person{ID: 1, Name: "John Doe", Email: "johndoe@example.com", Age: 30}

	// Insert the person into the store
	store.InsertPerson(person)

	// Retrieve the person by ID
	retrievedPerson, exists := store.GetPerson(1)
	if !exists {
		t.Fatalf("expected person with ID 1 to be found")
	}

	// Verify the retrieved person's data
	if retrievedPerson.Name != "John Doe" {
		t.Errorf("expected name 'John Doe', got %s", retrievedPerson.Name)
	}
	if retrievedPerson.Email != "johndoe@example.com" {
		t.Errorf("expected email 'johndoe@example.com', got %s", retrievedPerson.Email)
	}
	if retrievedPerson.Age != 30 {
		t.Errorf("expected age 30, got %d", retrievedPerson.Age)
	}
}
