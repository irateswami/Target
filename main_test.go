package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductEndpoint(t *testing.T) {

	var testGetBuffer = []byte(``)

	// Create a new request
	req, err := http.NewRequest("GET", "/product/123", bytes.NewBuffer(testGetBuffer))
	if err != nil {
		t.Fatal("new request for get test failed: ", err)
	}

	// Set the header so we know what we're dealing with
	req.Header.Set("content-type", "application/json")

	// Recorder to get the response
	rr := httptest.NewRecorder()

	// Handler is what's actually going to handle the call to the function
	handler := http.HandlerFunc(InsertProductEndpoint)

	// Tell the handler serve the http request and with what
	handler.ServeHTTP(rr, req)

	// See if the proper status code was returned
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("insert test returned a bad status code: GOT %v WANTED %v",
			status, http.StatusOK)
	}
}

func TestInsertDeleteProductEndpoint(t *testing.T) {

	// Byte test string for insert
	var testInsertBuffer = []byte(`{
		"productid":"123", 
		"productprice":"11.99",
		"productcurrency":"USD"
	}`)

	// Create a new insert request
	req, err := http.NewRequest("POST", "/product/123", bytes.NewBuffer(testInsertBuffer))
	if err != nil {
		t.Fatal("new request for insert test failed: ", err)
	}

	// Set the header so we know what we're dealing with
	req.Header.Set("content-type", "application/json")

	// Recorder to get the response
	rr := httptest.NewRecorder()

	// Handler is what's actually going to handle the call to the function
	handler := http.HandlerFunc(InsertProductEndpoint)

	// Tell the handler serve the http request and with what
	handler.ServeHTTP(rr, req)

	// See if the proper status code was returned
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("insert test returned a bad status code: GOT %v WANTED %v",
			status, http.StatusOK)
	}

	// Create a new delete request for what we just inserted
	req, err = http.NewRequest("DELETE", "/product/123", nil)
	if err != nil {
		t.Fatal("new request for delete test failed: ", err)
	}

	// Set the header again
	req.Header.Set("content-type", "application/json")

	// New recorder
	rr = httptest.NewRecorder()

	// New handler
	handler = http.HandlerFunc(DeleteProductEndpoint)

	// Serve it up once again
	handler.ServeHTTP(rr, req)
	if err != nil {
		t.Fatal("delete test returned a bad status code: ", err)
	}

}

func TestUpdateProductEndpoint(t *testing.T) {

	// Byte test string for upate
	var testUpdateBuffer = []byte(`{
		"productprice":"99.99"
	}`)

	// Create a new request
	req, err := http.NewRequest("PUT", "/product/123", bytes.NewBuffer(testUpdateBuffer))
	if err != nil {
		t.Fatal("new request for update test failed: ", err)
	}

	// Set the header so we know what we're dealing with
	req.Header.Set("content-type", "application/json")

	// Recorder to get the response
	rr := httptest.NewRecorder()

	// Handler is what's actually going to handle the call to the function
	handler := http.HandlerFunc(UpdateProductEndpoint)

	// Tell the handler serve the http request and with what
	handler.ServeHTTP(rr, req)

	// See if the proper status code was returned
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("update test returned a bad status code: GOT %v WANTED %v",
			status, http.StatusOK)
	}
}

func ProductNameEndpoint(t *testing.T) {

	// Id we're looking for
	id := "13860428"

	// Allocate a new product
	product := new(Product)

	// Hit the endpoint
	product.ProductNameEndpoint(id)

	// We definitely know what we're looking for
	expected := "The Big Lebowski (Blu-ray)"

	// Check if we got what we wanted
	if expected != product.Productname {
		t.Errorf("product name endpoint failed: GOT `%v`, WANTED `%v`",
			product.Productname, expected)
	}
}
