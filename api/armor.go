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

// Armor schema
type Armor struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string             `json:"name,omitempty" bson:"name,omitempty"`
	Rarity    int                `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Upgrade   int                `json:"upgrade,omitempty" bson:"upgrade,omitempty"`
	Level     int                `json:"level,omitempty" bson:"level,omitempty"`
	HeroLevel int                `json:"herolevel,omitempty" bson:"herolevel,omitempty"`
	Defense   []int              `json:"defense,omitempty" bson:"defense,omitempty"`
	Effects   []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}

// CreateArmorEndpoint creates armor on database
func CreateArmorEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var armor Armor
	json.NewDecoder(request.Body).Decode(&armor)
	collection := *client.Database(databaseName).Collection("armors")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, armor)
	json.NewEncoder(response).Encode(result)
}

// GetOneArmorEndpoint returns a json for one armor
func GetOneArmorEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var armor Armor

	collection := *client.Database(databaseName).Collection("armors")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Armor{ID: id}).Decode(&armor)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(armor)
}

// GetArmorEndpoint returns all armors
func GetArmorEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var armors []Armor
	collection := client.Database(databaseName).Collection("armors")
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
		var armor Armor
		cursor.Decode(&armor)
		armors = append(armors, armor)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(armors)
}

// DeleteArmorEndpoint deletes one armor based on it's id
func DeleteArmorEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("armors")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Armor{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
