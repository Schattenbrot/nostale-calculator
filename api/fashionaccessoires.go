package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Fashionaccessoire schema
type Fashionaccessoire struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Level     int                `json:"level,omitempty" bson:"level,omitempty"`
	HeroLevel int                `json:"herolevel,omitempty" bson:"herolevel,omitempty"`
	Defense   []int              `json:"defense,omitempty" bson:"defense,omitempty"`
	Effects   []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}

// CreateFashionaccessoireEndpoint creates fashionaccessoire on database
func CreateFashionaccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var fashionaccessoire Fashionaccessoire
	json.NewDecoder(request.Body).Decode(&fashionaccessoire)
	collection := *client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, fashionaccessoire)
	json.NewEncoder(response).Encode(result)
}

// GetOneFashionaccessoireEndpoint returns a json for one fashionaccessoire
func GetOneFashionaccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var fashionaccessoire Fashionaccessoire

	collection := *client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Fashionaccessoire{ID: id}).Decode(&fashionaccessoire)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(fashionaccessoire)
}

// GetFashionaccessoireEndpoint returns all fashionaccessoires
func GetFashionaccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var fashionaccessoires []Fashionaccessoire
	collection := client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var fashionaccessoire Fashionaccessoire
		cursor.Decode(&fashionaccessoire)
		fashionaccessoires = append(fashionaccessoires, fashionaccessoire)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(fashionaccessoires)
}

// GetFashionaccessoireTypeEndpoint returns all fashionaccessoires of a specified type
func GetFashionaccessoireTypeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	fashionaccessoiretype := params["fashionaccessoiretype"]
	var fashionaccessoires []Fashionaccessoire

	collection := client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, Fashionaccessoire{Type: fashionaccessoiretype})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var fashionaccessoire Fashionaccessoire
		cursor.Decode(&fashionaccessoire)
		fashionaccessoires = append(fashionaccessoires, fashionaccessoire)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(fashionaccessoires)
}

// DeleteFashionaccessoireEndpoint deletes one fashionaccessoire based on it's id
func DeleteFashionaccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Fashionaccessoire{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
