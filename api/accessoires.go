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

// Accessoire schema
type Accessoire struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type          string             `json:"type,omitempty" bson:"type,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Rarity        int                `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Upgrade       int                `json:"upgrade,omitempty" bson:"upgrade,omitempty"`
	Level         int                `json:"level,omitempty" bson:"level,omitempty"`
	HeroLevel     int                `json:"herolevel,omitempty" bson:"herolevel,omitempty"`
	Defense       []int              `json:"defense,omitempty" bson:"defense,omitempty"`
	Optioneffects []optioneffects    `json:"optioneffectss,omitempty" bson:"optioneffects,omitempty"`
	Effects       []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}

type optioneffects struct {
	Text  string `json:"text,omitempty" bson:"text,omitempty"`
	Value int    `json:"value,omitempty" bson:"value,omitempty"`
}

// CreateAccessoireEndpoint creates accessoire on database
func CreateAccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var accessoire Accessoire
	json.NewDecoder(request.Body).Decode(&accessoire)
	collection := *client.Database(databaseName).Collection("accessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, accessoire)
	json.NewEncoder(response).Encode(result)
}

// GetOneAccessoireEndpoint returns a json for one accessoire
func GetOneAccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var accessoire Accessoire

	collection := *client.Database(databaseName).Collection("accessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Accessoire{ID: id}).Decode(&accessoire)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(accessoire)
}

// GetAccessoireEndpoint returns all accessoires
func GetAccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var accessoires []Accessoire
	collection := client.Database(databaseName).Collection("accessoires")
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
		var accessoire Accessoire
		cursor.Decode(&accessoire)
		accessoires = append(accessoires, accessoire)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(accessoires)
}

// GetAccessoireTypeEndpoint returns all accessoires of a specified type
func GetAccessoireTypeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	accessoiretype := params["accessoiretype"]
	var accessoires []Accessoire

	collection := client.Database(databaseName).Collection("accessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, Accessoire{Type: accessoiretype})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var accessoire Accessoire
		cursor.Decode(&accessoire)
		accessoires = append(accessoires, accessoire)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(accessoires)
}

// DeleteAccessoireEndpoint deletes one accessoire based on it's id
func DeleteAccessoireEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("accessoires")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Accessoire{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
