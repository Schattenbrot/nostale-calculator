package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var err error
var databaseName string = "nostaleDB"

func main() {
	fmt.Println("Starting the application...")

	// Make and Connect new Client
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	// Router
	mainRouter := mux.NewRouter()
	weaponSubrouter := mainRouter.PathPrefix("/api/weapon").Subrouter()
	SetWeaponSubRouter(weaponSubrouter)
	armorSubrouter := mainRouter.PathPrefix("/api/armor").Subrouter()
	SetArmorSubRouter(armorSubrouter)
	fairySubrouter := mainRouter.PathPrefix("/api/fairy").Subrouter()
	SetFairySubRouter(fairySubrouter)
	fashionaccessoireSubrouter := mainRouter.PathPrefix("/api/fashionaccessoire").Subrouter()
	SetFashionaccessoireSubRouter(fashionaccessoireSubrouter)
	resistanceSubrouter := mainRouter.PathPrefix("/api/resistance").Subrouter()
	SetResistanceSubRouter(resistanceSubrouter)
	accessoireSubrouter := mainRouter.PathPrefix("/api/accessoire").Subrouter()
	SetAccessoireSubRouter(accessoireSubrouter)
	costumeSubrouter := mainRouter.PathPrefix("/api/costume").Subrouter()
	SetCostumeSubRouter(costumeSubrouter)

	handler := cors.Default().Handler(mainRouter)
	http.ListenAndServe(":3000", handler)
}
