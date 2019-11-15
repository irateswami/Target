package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Because we have one example and I'm not totally sure on how the api url string is constructed, this will just be a constant. In the real world this would be constructed via arguments.
var redskyURL string

// Mongo db client is going to be use a lot of places, and there's only one db we're hitting, so make it gobal
var client *mongo.Client

func init() {

	redskyURL = "https://redsky.target.com/v2/pdp/tcin/13860428?excludes=taxonomy,price,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics"

	var err error

	// Get our api address to hit, here we're using a local mongodb docker container
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("db didn't connect: ", err)
	}
}

type Product struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Productid       string             `json:"productid,omitempty" bson:"productid,omitempty"`
	Productname     string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Productprice    string             `json:"productprice,omitempty" bson:"productprice,omitempty"`
	ProductCurrency string             `json:"productcurrency,omitempty" bson:"productcurrency,omitempty"`
}

func (p *Product) ProductNameEndpoint(id string) {
	fmt.Println("ProductName called")

	// Construct http request to the product info api. In this case the api address is a string constant, but in the real world this would be constructed via another function
	response, err := http.Get(redskyURL)
	if err != nil {
		log.Fatal("product information api request failed: ", err)
	}
	defer response.Body.Close()

	// Read in the data using io
	unstringData, err := ioutil.ReadAll(response.Body)
	if err != nil {

		log.Fatal("ioutil readall failed: ", err)
	}

	// Stringify the data
	data := string(unstringData)

	// Big thanks to Josh Baker for making this json parser with syntax that actually makes sense
	productName := gjson.Get(data, "product.item.product_description.title").String()

	p.Productname = string(productName)

}

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
	product.ProductNameEndpoint(redskyURL)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Encode the response onto the product
	json.NewEncoder(response).Encode(product)

}

func InsertProductEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Print("InsertPoduct called")

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

	update := bson.M{"$set": bson.M{"productprice": price}}

	// Establish the db and collection we're going to use
	collection := client.Database("target").Collection("products")

	// Set a timeout, defer the timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//
	_, err := collection.UpdateOne(ctx, Product{Productid: id}, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
}

func main() {

	// Create a router through which traffic will flow
	router := mux.NewRouter()

	// Give it some places to go
	router.HandleFunc("/product/{id}", InsertProductEndpoint).Methods("POST")
	router.HandleFunc("/product/{id}", GetProductEndpoint).Methods("GET")
	router.HandleFunc("/product/{id}", UpdateProductEndpoint).Methods("PUT")

	//Serve up everything on port 8080
	http.ListenAndServe(":8080", router)

	// Create a timeout, defer the cancel
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Defer closing until we exit
	defer client.Disconnect(ctx)
}
