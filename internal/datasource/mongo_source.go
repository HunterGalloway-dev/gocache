package datasource

import (
	"context"
	"fmt"
	"gocache/pkg/model"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoSource struct {
	db *mongo.Client
}

var (
	username = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_ROOT_PASSWORD")
	host     = os.Getenv("DB_HOST")
	port     = os.Getenv("DB_PORT")
)

func NewMongo() DataSource {
	fmt.Println(username, password, host, port)
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", username, password, host, port)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))

	if err != nil {
		log.Fatal(err)

	}
	return &mongoSource{
		db: client,
	}
}

func (s *mongoSource) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.db.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

// Person methods
func (s *mongoSource) GetAllPersons() ([]model.Person, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Printf("DATASOURCE: GetAllPersons called")

	collection := s.db.Database("gocache").Collection("person")

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("DATASOURCE: GetAllPersons error getting collection: %v", err)
		return nil, err
	}

	var persons []model.Person
	if err = cursor.All(ctx, &persons); err != nil {
		log.Printf("DATASOURCE: GetAllPersons error finding all on collection: %v", err)
		return nil, err
	}

	log.Printf("DATASOURCE: GetAllPersons success: found %v persons", len(persons))

	return persons, nil
}