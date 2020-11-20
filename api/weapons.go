package main

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
	"fmt"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Weapon schema
type Weapon struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Weapontype    string             `json:"weapontype,omitempty" bson:"weapontype,omitempty"`
	Name          string             `json:"name,omitempty" bson:"name,omitempty"`
	Rarity        int                `json:"rarity,omitempty" bson:"rarity,omitempty"`
	Upgrade       int                `json:"upgrade,omitempty" bson:"upgrade,omitempty"`
	Level         int                `json:"level,omitempty" bson:"level,omitempty"`
	HeroLevel     int                `json:"herolevel,omitempty" bson:"herolevel,omitempty"`
	Damage        int                `json:"damage,omitempty" bson:"damage,omitempty"`
	Concentration int                `json:"concentration,omitempty" bson:"concentration,omitempty"`
	CritDamage    float32            `json:"critDamage,omitempty" bson:"critDamage,omitempty"`
	CritChance    float32            `json:"critChance,omitempty" bson:"critChance,omitempty"`
	Effects       []Effect           `json:"effects,omitempty" bson:"effects,omitempty"`
}

// CreateWeaponEndpoint creates weapon on database
func CreateWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("CreateWeaponEndpoint used...")
	response.Header().Add("content-type", "application/json")
	var weapon Weapon
	json.NewDecoder(request.Body).Decode(&weapon)
	collection := *client.Database(databaseName).Collection("weapons")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, weapon)
	json.NewEncoder(response).Encode(result)
}

// GetOneWeaponEndpoint returns a json for one weapon
func GetOneWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var weapon Weapon

	collection := *client.Database(databaseName).Collection("weapons")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, Weapon{ID: id}).Decode(&weapon)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(weapon)
}

// GetWeaponEndpoint returns all weapons
func GetWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var weapons []Weapon
	collection := client.Database(databaseName).Collection("weapons")
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
		var weapon Weapon
		cursor.Decode(&weapon)
		weapons = append(weapons, weapon)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(weapons)
}

// GetWeaponTypeEndpoint returns all weapons of a specified type
func GetWeaponTypeEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	weapontype := params["weapontype"]
	var weapons []Weapon

	collection := client.Database(databaseName).Collection("weapons")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, Weapon{Weapontype: weapontype})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var weapon Weapon
		cursor.Decode(&weapon)
		weapons = append(weapons, weapon)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(weapons)
}

// DeleteWeaponEndpoint deletes one weapon based on it's id
func DeleteWeaponEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := client.Database(databaseName).Collection("weapons")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Delete().SetCollation(&options.Collation{
		Locale:    "en_US",
		Strength:  1,
		CaseLevel: false,
	})

	_, err := collection.DeleteOne(ctx, Weapon{ID: id}, opts)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(`{ message: "OK" }`))
}
