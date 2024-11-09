package services

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId    primitive.ObjectID `json:"user_id,omitempty" bson:"user_id,omitempty"`
	Submitted bool               `json:"submitted" bson:"submitted"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

func (s *Submission) CreateSubmission(entry Submission) (Submission, error) {
	collection := ReturnCollectPointer("submission")

	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	filer := bson.D{{Key: "_id", Value: entry.ID}}

	var user User

	err := collection.FindOne(context.TODO(), filer).Decode(&user)

	if err == nil {
		insertResult, err := collection.InsertOne(context.Background(), entry)

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
