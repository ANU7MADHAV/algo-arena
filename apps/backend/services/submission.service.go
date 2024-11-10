package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Submission struct {
	ID         primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	UserId     string             `json:"user_id,omitempty" bson:"user_id,omitempty"`
	LanguageID int                `json:"language_id,omitempty" bson:"language_id,omitempty"`
	SourceCode string             `json:"source_code,omitempty" bson:"source_code,omitempty"`
	Stdin      string             `json:"stdin,omitempty" bson:"stdin,omitempty"`
	Judge0     string             `json:"judge0,omitempty" bson:"judge0,omitempty"`
	Submitted  bool               `json:"submitted" bson:"submitted"`
	CreatedAt  time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}

type Judge0 struct {
	Token string `json:"token"`
}

type Response struct {
	Stdout string `json:"stdout"`
}

func (s Submission) CreateSubmission(entry Submission) (Submission, error) {
	var user User
	collection := ReturnCollectPointer("submission")
	url := "https://judge0-ce.p.rapidapi.com/submissions?base64_encoded=true&wait=false&fields=*"

	languageID := entry.LanguageID
	sourceCode := entry.SourceCode
	stdin := entry.Stdin

	payload := fmt.Sprintf(`{"language_id":%d,"source_code":"%s","stdin":"%s"}`, languageID, sourceCode, stdin)
	reader := strings.NewReader(payload)

	req, _ := http.NewRequest("POST", url, reader)

	req.Header.Add("x-rapidapi-key", "58facff32fmsh848026aeeec17cfp19c063jsn3dd771b83112")
	req.Header.Add("x-rapidapi-host", "judge0-ce.p.rapidapi.com")
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
	fmt.Println("body", string(body))

	entry.CreatedAt = time.Now()
	entry.UpdatedAt = time.Now()

	er := user.GetUserById(entry.UserId)

	fmt.Println("err", er)

	token := string(body)

	fmt.Println("helo", token)

	var judge0 Judge0

	checkError := json.Unmarshal([]byte(token), &judge0)

	if checkError != nil {
		log.Fatal("Error parsing Judge0 JSON:", err)
	}
	//
	fmt.Println("Token:", judge0.Token)

	GetSubmission(judge0.Token)

	entry.Judge0 = judge0.Token

	if err == nil {
		insertResult, err := collection.InsertOne(context.Background(), entry)
		//
		// 		fmt.Println("result", insertResult)

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

func GetSubmission(token string) (std string, err error) {

	url := fmt.Sprintf("https://judge0-ce.p.rapidapi.com/submissions/%s?_encoded=true&fields=*", token)

	fmt.Println("hitted get")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	req.Header.Add("x-rapidapi-key", "58facff32fmsh848026aeeec17cfp19c063jsn3dd771b83112")
	req.Header.Add("x-rapidapi-host", "judge0-ce.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// fmt.Println("body get", body)

	respons := []byte(body)

	var response Response

	if err := json.Unmarshal(respons, &response); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Stdout:", response.Stdout)
	//
	// 	fmt.Println("res", res)
	// 	fmt.Println("body", string(body))

	return response.Stdout, nil
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
