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

// Resistance schema
type Resistance struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Type      string             `json:"type,omitempty" bson:"type,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Level     int                `json:"level,omitempty" bson:"level,omitempty"`
	HeroLevel int                `json:"herolevel,omitempty" bson:"herolevel,omitempty"`
	Defense   []int              `json:"defense,omitempty" bson:"defense,omitempty"`
	Effects   []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}

// CreateResistanceEndpoint creates resistance on database
func CreateResistanceEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var resistance Resistance
	json.NewDecoder(request.Body).Decode(&resistance)
	collection := *client.Database(databaseName).Collection("resistances")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, resistance)
	json.NewEncoder(response).Encode(result)
}

// GetOneResistanceEndpoint returns a json for one resistance
func GetOneResistanceEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var resistance Resistance

	collection := *client.Database(databaseName).Collection("resistances")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Resistance{ID: id}).Decode(&resistance)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(resistance)
}

// GetResistanceEndpoint returns all resistances
func GetResistanceEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var resistances []Resistance
	collection := client.Database(databaseName).Collection("resistances")
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
		var resistance Resistance
		cursor.Decode(&resistance)
		resistances = append(resistances, resistance)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(resistances)
}

// GetResistanceTypeEndpoint returns all resistances of a specified type
func GetResistanceTypeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	resistancetype := params["resistancetype"]
	var resistances []Resistance

	collection := client.Database(databaseName).Collection("resistances")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, Resistance{Type: resistancetype})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var resistance Resistance
		cursor.Decode(&resistance)
		resistances = append(resistances, resistance)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(resistances)
}

// DeleteResistanceEndpoint deletes one resistance based on it's id
func DeleteResistanceEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("resistances")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Resistance{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
