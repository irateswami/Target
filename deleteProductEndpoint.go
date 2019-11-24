package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteProductEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("DeleteProduct called")

	// Set the header
	response.Header().Set("content-type", "application/json")

	// Grab the params passed
	params := mux.Vars(request)

	// Grab the id from the params
	id := params["id"]

	// Define our db collection
	collection := client.Database("target").Collection("products")

	// Set a timeout and defer
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Try to delete, error out if something went wrong
	_, err := collection.DeleteOne(ctx, bson.M{"productid": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message: " ` + err.Error() + `" }`))
	}
}
