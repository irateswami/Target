package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	configs Configs
	client  *mongo.Client
)

func init() {

	// Grab the configs from the file
	configsFile, err := os.Open("configs.json")
	defer configsFile.Close()
	if err != nil {
		log.Fatal("configs weren't found: ", err)
	}

	// Read in the byte string using the io utility
	bytesFileIn, err := ioutil.ReadAll(configsFile)
	if err != nil {
		log.Fatal("config file wasn't read in properly: ", err)
	}

	// Unmarshall the byte string into the configs object address
	json.Unmarshal(bytesFileIn, &configs)

	// Get our api address to hit, here we're using a local mongodb docker container
	clientOptions := options.Client().ApplyURI(configs.DBURL)

	// Try to connect and if failure, let me know
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("db didn't connect: ", err)
	}
}

func main() {

	// Create a router through which traffic will flow
	router := mux.NewRouter()

	// Give it some places to go
	router.HandleFunc("/product/{id}", InsertProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/product/{id}", UpdateProductEndpoint).Methods("PUT")
	router.HandleFunc("/product/{id}", DeleteProductEndpoint).Methods("DELETE")

	// Serve up everything on port 8080
	http.ListenAndServe(":8080", router)

	// Create a timeout, defer the cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Defer closing until we exit
	defer client.Disconnect(ctx)
}
