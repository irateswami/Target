package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func InsertProductEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("InsertPoduct called")

	// Set the header
	response.Header().Set("content-type", "application/json")

	// Allocate a new Product
	product := new(Product)

	// Provide a new decoder for the product we just allocated
	json.NewDecoder(request.Body).Decode(&product)

	// Establish the db and collection we're going to use
	collection := client.Database("target").Collection("products")

	// Set a timeout, defer the timeout
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	// Grab everything and try inserting into the database
	result, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Fatal("insert failed: ", err)
	}

	// Encode the result
	json.NewEncoder(response).Encode(result)
}
