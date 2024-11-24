package store

import (
	"gocache/pkg/model"
	"testing"
)

func TestInsertPerson(t *testing.T) {
	store := NewKVStore()

	p := model.Person{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}

	store.InsertPerson(p)

	if person, ok := store.GetPerson(1); !ok || person != p {
		t.Errorf("expected person %+v, got %+v", p, person)
	}
}

func TestInsertPersons(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	store.InsertPersons(persons)

	for _, p := range persons {
		if person, ok := store.GetPerson(p.ID); !ok || person != p {
			t.Errorf("expected person %+v, got %+v", p, person)
		}
	}
}

func TestQuery(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com", Age: 30},
	}

	store.InsertPersons(persons)

	result := store.Query("", "jane@example.com", nil)
	if len(result) != 1 {
		t.Errorf("expected to find Jane Smith, got %+v", result)
	}

	result = store.Query("", "", []int{30})
	if len(result) != 2 {
		t.Errorf("expected 2 persons with age 30, got %+v", result)
	}
}

func TestDeletePerson(t *testing.T) {
	store := NewKVStore()

	p := model.Person{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}

	store.InsertPerson(p)

	if _, ok := store.GetPerson(1); !ok {
		t.Fatal("expected person to be in the store before deletion")
	}

	if err := store.DeletePerson(1); err != nil {
		t.Fatalf("unexpected error deleting person: %v", err)
	}

	if _, ok := store.GetPerson(1); ok {
		t.Fatal("expected person to be deleted from the store")
	}
}

func TestDeletePersonNotFound(t *testing.T) {
	store := NewKVStore()

	err := store.DeletePerson(999)
	if err == nil || err.Error() != "person not found" {
		t.Errorf("expected 'person not found' error, got %v", err)
	}
}

func TestUpdatePerson(t *testing.T) {
	store := NewKVStore()

	p := model.Person{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}

	store.InsertPerson(p)

	updatedPerson := model.Person{ID: 1, Name: "John Doe", Email: "john.doe@newdomain.com", Age: 31}
	if err := store.UpdatePerson(1, updatedPerson); err != nil {
		t.Fatalf("unexpected error updating person: %v", err)
	}

	person, ok := store.GetPerson(1)
	if !ok || person.Email != updatedPerson.Email || person.Age != updatedPerson.Age {
		t.Errorf("expected updated person %+v, got %+v", updatedPerson, person)
	}
}

func TestUpdatePersonNotFound(t *testing.T) {
	store := NewKVStore()

	err := store.UpdatePerson(999, model.Person{ID: 999, Name: "Non-existent", Email: "nonexistent@example.com"})
	if err == nil || err.Error() != "person not found" {
		t.Errorf("expected 'person not found' error, got %v", err)
	}
}

func TestGetAllPersons(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	store.InsertPersons(persons)

	result := store.GetAllPersons()
	if len(result) != 2 {
		t.Errorf("expected 2 persons, got %d", len(result))
	}
}

func TestInsertPersonTwice(t *testing.T) {
	store := NewKVStore()

	p := model.Person{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}
	store.InsertPerson(p)

	store.InsertPerson(p)

	if person, ok := store.GetPerson(1); !ok || person != p {
		t.Errorf("expected person %+v, got %+v", p, person)
	}
}

func TestQueryWithName(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "John Doe", Email: "john.doe@example.com", Age: 35},
	}

	store.InsertPersons(persons)

	result := store.Query("John Doe", "", nil)
	if len(result) != 2 {
		t.Errorf("expected 2 persons with name 'John Doe', got %d", len(result))
	}

	for _, person := range result {
		if person.Name != "John Doe" {
			t.Errorf("expected name 'John Doe', got '%s'", person.Name)
		}
	}
}

func TestQueryWithEmail(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "John Doe", Email: "john.doe@example.com", Age: 35},
	}

	store.InsertPersons(persons)

	result := store.Query("", "john.doe@example.com", nil)
	if len(result) != 1 || result[0].Email != "john.doe@example.com" {
		t.Errorf("expected 1 person with email 'john.doe@example.com', got %+v", result)
	}
}

func TestQueryWithAge(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com", Age: 30},
	}

	store.InsertPersons(persons)

	result := store.Query("", "", []int{30})
	if len(result) != 2 {
		t.Errorf("expected 2 persons with age 30, got %d", len(result))
	}

	for _, person := range result {
		if person.Age != 30 {
			t.Errorf("expected age 30, got %d", person.Age)
		}
	}
}

func TestQueryWithEmailAndAge(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com", Age: 30},
		{ID: 4, Name: "John Doe", Email: "john.doe@example.com", Age: 30},
	}

	store.InsertPersons(persons)

	result := store.Query("", "john@example.com", []int{30})
	if len(result) != 1 || result[0].Email != "john@example.com" {
		t.Errorf("expected 1 person with email 'john@example.com' and age 30, got %+v", result)
	}
}

func TestQueryWithNameAndAge(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com", Age: 30},
		{ID: 4, Name: "John Doe", Email: "john.doe@example.com", Age: 35},
	}

	store.InsertPersons(persons)

	result := store.Query("John Doe", "", []int{30})
	if len(result) != 1 || result[0].Name != "John Doe" || result[0].Age != 30 {
		t.Errorf("expected 1 person with name 'John Doe' and age 30, got %+v", result)
	}
}

func TestQueryWithAllFields(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
		{ID: 3, Name: "Alice Johnson", Email: "alice@example.com", Age: 30},
		{ID: 4, Name: "John Doe", Email: "john.doe@example.com", Age: 30},
	}

	store.InsertPersons(persons)

	result := store.Query("John Doe", "john@example.com", []int{30})
	if len(result) != 1 || result[0].Email != "john@example.com" || result[0].Name != "John Doe" {
		t.Errorf("expected 1 person with name 'John Doe', email 'john@example.com', and age 30, got %+v", result)
	}
}

func TestQueryWithNoResults(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	store.InsertPersons(persons)

	result := store.Query("Non Existent", "nonexistent@example.com", []int{99})
	if len(result) != 0 {
		t.Errorf("expected no results, got %+v", result)
	}
}

func TestQueryWithEmptyFields(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	store.InsertPersons(persons)

	result := store.Query("", "", nil)
	if len(result) != 2 {
		t.Errorf("expected 2 persons, got %d", len(result))
	}
}

func TestQueryWithNoCriteria(t *testing.T) {
	store := NewKVStore()

	persons := []model.Person{
		{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30},
		{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Age: 25},
	}

	store.InsertPersons(persons)

	result := store.Query("", "", nil)
	if len(result) != 2 {
		t.Errorf("expected 2 persons, got %d", len(result))
	}
}
