package model

type Person struct {
	ID    int    `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Age   int    `json:"age" bson:"age"`
	Email string `json:"email" bson:"email"`
}
