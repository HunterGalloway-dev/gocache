package datasource

import (
	"context"
	"fmt"
	"gocache/internal/logger"
	"gocache/pkg/model"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoSource struct {
	db         *mongo.Client
	personColl *mongo.Collection
}

var (
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_ROOT_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
	name     = os.Getenv("DB_NAME")
	coll     = os.Getenv("COLLECTION_NAME")
)

func NewMongo() (DataSource, error) {
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", username, password, host, port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	personColl := client.Database(name).Collection(coll)

	if err != nil {
		logger.Logger.Fatalf("Error connecting to MongoDB: %v", err)
		return nil, err

	}
	return &mongoSource{
		db:         client,
		personColl: personColl,
	}, nil
}

func (m *mongoSource) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := m.db.Ping(ctx, nil)
	if err != nil {
		logger.Logger.Errorf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

// Person methods
func (m *mongoSource) GetAllPersons() ([]model.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Logger.Info("DATASOURCE: GetAllPersons called")

	cursor, err := m.personColl.Find(ctx, bson.D{})
	if err != nil {
		logger.Logger.Errorf("DATASOURCE: GetAllPersons error getting collection: %v", err)
		return nil, err
	}

	var persons []model.Person
	if err = cursor.All(ctx, &persons); err != nil {
		logger.Logger.Errorf("DATASOURCE: GetAllPersons error finding all on collection: %v", err)
		return nil, err
	}

	logger.Logger.Infof("DATASOURCE: GetAllPersons success: found %v persons", len(persons))

	return persons, nil
}

func (m *mongoSource) UpdatePerson(person model.Person) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	logger.Logger.Infof("DATASOURCE: UpdatePerson called")

	filter := bson.D{{Key: "id", Value: person.ID}}
	update := bson.D{{Key: "$set", Value: person}}

	_, err := m.personColl.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Logger.Errorf("DATASOURCE: UpdatePerson error updating person: %v", err)
		return err
	}

	logger.Logger.Infof("DATASOURCE: UpdatePerson success: updated person with ID %v", person.ID)

	return nil
}
