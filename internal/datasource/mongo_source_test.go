package datasource

import (
	"context"
	"gocache/pkg/model"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/bson"
)

func mustStartMongoContainer() (func(context.Context) error, error) {
	ctx := context.Background()
	dbContainer, err := mongodb.Run(ctx, "mongo:latest", mongodb.WithUsername("root"), mongodb.WithPassword("password"), testcontainers.WithWaitStrategy(wait.ForLog("Waiting for connections")))
	if err != nil {
		return nil, err
	}

	dbHost, err := dbContainer.Host(ctx)
	if err != nil {
		return dbContainer.Terminate, err
	}

	dbPort, err := dbContainer.MappedPort(ctx, "27017/tcp")
	if err != nil {
		return dbContainer.Terminate, err
	}

	// Set the connection details
	host = dbHost
	port = dbPort.Port()
	username = "root"
	password = "password"
	name = "gocache"
	coll = "person"

	return dbContainer.Terminate, nil
}

func TestMain(m *testing.M) {
	mustStartMongoContainer()
	m.Run()
}

func TestNew(t *testing.T) {
	mongo, err := NewMongo()
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	if mongo == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	mongo, err := NewMongo()
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	health := mongo.Health()
	if health["message"] != "It's healthy" {
		t.Fatal("Health() returned unhealthy")
	}
}
func TestGetAllPersons(t *testing.T) {
	terminate, err := mustStartMongoContainer()
	if err != nil {
		t.Fatalf("Failed to start MongoDB container: %v", err)
	}
	defer terminate(context.Background())

	mongo, err := NewMongo()
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	// Insert test data
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.(*mongoSource).db.Database("gocache").Collection("person")
	_, err = collection.InsertMany(ctx, []interface{}{
		bson.D{{Key: "name", Value: "John Doe"}, {Key: "age", Value: 30}},
		bson.D{{Key: "name", Value: "Jane Doe"}, {Key: "age", Value: 25}},
	})
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Test GetAllPersons
	persons, err := mongo.GetAllPersons()
	if err != nil {
		t.Fatalf("GetAllPersons() error: %v", err)
	}

	if len(persons) != 2 {
		t.Fatalf("Expected 2 persons, got %d", len(persons))
	}

	expectedNames := []string{"John Doe", "Jane Doe"}
	for i, person := range persons {
		if person.Name != expectedNames[i] {
			t.Errorf("Expected name %s, got %s", expectedNames[i], person.Name)
		}
	}
}

func TestUpdatePerson(t *testing.T) {
	terminate, err := mustStartMongoContainer()
	if err != nil {
		t.Fatalf("Failed to start MongoDB container: %v", err)
	}
	defer terminate(context.Background())

	mongo, err := NewMongo()
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}

	// Insert test data
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := mongo.(*mongoSource).db.Database("gocache").Collection("person")
	_, err = collection.InsertOne(ctx, bson.D{{Key: "id", Value: 1}, {Key: "name", Value: "John Doe"}, {Key: "age", Value: 30}, {Key: "email", Value: "john.doe@example.com"}})
	if err != nil {
		t.Fatalf("Failed to insert test data: %v", err)
	}

	// Update the person
	person := model.Person{ID: 1, Name: "John Smith", Age: 35, Email: "john.smith@example.com"}
	err = mongo.UpdatePerson(person)
	if err != nil {
		t.Fatalf("UpdatePerson() error: %v", err)
	}

	// Verify the update
	var updatedPerson model.Person
	err = collection.FindOne(ctx, bson.D{{Key: "id", Value: person.ID}}).Decode(&updatedPerson)
	if err != nil {
		t.Fatalf("Failed to find updated person: %v", err)
	}

	if updatedPerson.Name != person.Name {
		t.Errorf("Expected name %s, got %s", person.Name, updatedPerson.Name)
	}

	if updatedPerson.Age != person.Age {
		t.Errorf("Expected age %d, got %d", person.Age, updatedPerson.Age)
	}

	if updatedPerson.Email != person.Email {
		t.Errorf("Expected email %s, got %s", person.Email, updatedPerson.Email)
	}
}
