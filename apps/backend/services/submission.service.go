package services

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Submitted bool               `json:"submitted" bson:"submitted"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (s Submission) CreateSubmission(entry Submission) (Submission, error) {
	var user User
	collection := ReturnCollectPointer("submission")

	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	err := user.GetUserById(entry.UserId)

	fmt.Println("err", err)

	if err == nil {
		insertResult, err := collection.InsertOne(context.Background(), entry)

		fmt.Println("result", insertResult)

		if err != nil {
			panic(err)
		}

		entry.ID = insertResult.InsertedID.(primitive.ObjectID)
	}
	if err != nil {
		panic(err)
	}
	return entry, nil

}

func (s *Submission) GetSubmissionById(id string) (Submission, error) {
	collection := ReturnCollectPointer("submission")

	mongoId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{Key: "_id", Value: mongoId}}

	var submission Submission

	if err = collection.FindOne(context.TODO(), filter).Decode(&submission); err != nil {
		log.Fatal(err)
		panic(err)
	}
	return submission, nil
}

func (s *Submission) GetAllSubmissions() ([]Submission, error) {
	collection := ReturnCollectPointer("submission")

	// fmt.Println("collections", collection)

	var submissions []Submission

	cursor, err := collection.Find(context.Background(), bson.D{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var submission Submission

		if err = cursor.Decode(&submission); err != nil {
			log.Fatal(err)
		}
		fmt.Println("sub", submission)
		submissions = append(submissions, submission)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
		return submissions, err
	}

	return submissions, nil
}
