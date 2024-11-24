package controller

import (
	"gocache/internal/datasource"
	"testing"
)

func TestPersonControllerNew(t *testing.T) {
	// Test the NewPersonController function
	db := datasource.NewMockDataSource()
	pc, err := NewPersonController(db)

	if err != nil {
		t.Fatalf("NewPersonController() returned an error: %v", err)
	}

	if pc == nil {
		t.Fatal("NewPersonController() returned nil")
	}
}

func TestPersonControllerHealth(t *testing.T) {
	// Test the Health function
	db := datasource.NewMockDataSource()
	pc, _ := NewPersonController(db)
	health := pc.Health()

	if health["status"] != "healthy" {
		t.Fatalf("Health() returned unhealthy: %v", health)
	}
}

func TestPersonControllerQuery(t *testing.T) {
	// Test the Query function
	db := datasource.NewMockDataSource()
	pc, _ := NewPersonController(db)
	persons, err := pc.Query("", "", nil)

	if err != nil {
		t.Fatalf("Query() returned an error: %v", err)
	}

	if len(persons) != 2 {
		t.Fatalf("Expected 2 persons, got %d", len(persons))
	}
}

func TestPersonControllerGetAllPersons(t *testing.T) {
	// Test the GetAllPersons function
	db := datasource.NewMockDataSource()
	pc, _ := NewPersonController(db)
	persons, err := pc.GetAllPersons()

	if err != nil {
		t.Fatalf("GetAllPersons() returned an error: %v", err)
	}

	if len(persons) != 2 {
		t.Fatalf("Expected 2 persons, got %d", len(persons))
	}
}

// Test the Query function with a name filter
func TestPersonControllerQueryWithName(t *testing.T) {
	db := datasource.NewMockDataSource()
	pc, _ := NewPersonController(db)
	persons, err := pc.Query("John Doe", "", nil)

	if err != nil {
		t.Fatalf("Query() returned an error: %v", err)
	}

	if len(persons) != 1 {
		t.Fatalf("Expected 1 person, got %d", len(persons))
	}

	if persons[0].Name != "John Doe" {
		t.Fatalf("Expected name John Doe, got %s", persons[0].Name)
	}
}
