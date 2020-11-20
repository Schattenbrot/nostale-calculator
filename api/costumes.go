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

// Costume schema
type Costume struct {
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type           string             `json:"type,omitempty" bson:"type,omitempty"`
	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
	Defense        []int              `json:"defense,omitempty" bson:"defense,omitempty"`
	CostumeEffects []costumeEffects   `json:"costumeeffects,omitempty" bson:"costumeeffects,omitempty"`
	Effects        []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}
type costumeEffects struct {
	Text  string `json:"text,omitempty" bson:"text,omitempty"`
	Value int    `json:"value,omitempty" bson:"value,omitempty"`
}

// CreateCostumeEndpoint creates costume on database
func CreateCostumeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var costume Costume
	json.NewDecoder(request.Body).Decode(&costume)
	collection := *client.Database(databaseName).Collection("costumes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, costume)
	json.NewEncoder(response).Encode(result)
}

// GetOneCostumeEndpoint returns a json for one costume
func GetOneCostumeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var costume Costume

	collection := *client.Database(databaseName).Collection("costumes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Costume{ID: id}).Decode(&costume)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(costume)
}

// GetCostumeEndpoint returns all costumes
func GetCostumeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var costumes []Costume
	collection := client.Database(databaseName).Collection("costumes")
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
		var costume Costume
		cursor.Decode(&costume)
		costumes = append(costumes, costume)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(costumes)
}

// GetCostumeTypeEndpoint returns all fashionaccessoires of a specified type
func GetCostumeTypeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	costume := params["costume"]
	var fashionaccessoires []Costume

	collection := client.Database(databaseName).Collection("fashionaccessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, Costume{Type: costume})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var fashionaccessoire Costume
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

// DeleteCostumeEndpoint deletes one costume based on it's id
func DeleteCostumeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("costumes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Costume{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
