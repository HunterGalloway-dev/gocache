package datasource

import (
	"context"
	"testing"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
	"github.com/testcontainers/testcontainers-go/wait"
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

	return dbContainer.Terminate, nil
}

func TestMain(m *testing.M) {
	mustStartMongoContainer()
	m.Run()
}

func TestNew(t *testing.T) {
	mongo := NewMongo()
	if mongo == nil {
		t.Fatal("New() returned nil")
	}
}

func TestHealth(t *testing.T) {
	mongo := NewMongo()
	health := mongo.Health()
	if health["message"] != "It's healthy" {
		t.Fatal("Health() returned unhealthy")
	}
}
