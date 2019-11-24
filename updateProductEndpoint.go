package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateProductEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("UpdateProduct called")

	// Set the header
	response.Header().Set("content-type", "application/json")

	// Grab the request parameters
	params := mux.Vars(request)

	// Grab the id from the request map
	id := params["id"]

	// Allocate a new product
	product := new(Product)

	// Decode the request body and assign it to product
	json.NewDecoder(request.Body).Decode(&product)

	// Grab the new price
	price := product.Productprice

	// Do some bson magic
	update := bson.M{"$set": bson.M{"productprice": price}}

	// Establish the db and collection we're going to use
	collection := client.Database("target").Collection("products")

	// Set a timeout, defer the timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Update the product, error out if something was amiss
	_, err := collection.UpdateOne(ctx, Product{Productid: id}, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
}
