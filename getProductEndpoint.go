package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetProductEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("GetProduct called")

	// Set the header
	response.Header().Set("content-type", "application/json")

	// Grab the request parameters
	params := mux.Vars(request)

	// Grab the id from the request map
	id := params["id"]

	// Allocate a new product
	product := new(Product)

	// Establish the db and collection we're going to use
	collection := client.Database("target").Collection("products")

	// Establish an acceptable timeout, defer the cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Try finding the product by it's product id
	err := collection.FindOne(ctx, Product{Productid: id}).Decode(&product)
	product.ProductNameExternal(id)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Encode the response onto the product
	json.NewEncoder(response).Encode(product)

}
