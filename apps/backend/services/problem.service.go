package services

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Problem struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description" bson:"description"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (p *Problem) CreateProblem(entry Problem) (Problem, error) {
	collection := ReturnCollectPointer("problem")

	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	insertResult, err := collection.InsertOne(context.TODO(), entry)

	if err != nil {
		log.Println("Error:", err)
		return Problem{}, err
	}

	entry.ID = insertResult.InsertedID.(primitive.ObjectID)

	return entry, nil
}

func (p *Problem) ListProblems() ([]Problem, error) {
	collction := ReturnCollectPointer("problem")

	var problems []Problem

	cursor, err := collction.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var problem Problem

		cursor.Decode(&problem)
		problems = append(problems, problem)
	}

	return problems, nil
}
