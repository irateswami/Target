package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetProductEndpoint(t *testing.T) {

}

func TestInsertProductEndpoint(t *testing.T) {

	// Byte test string for insert
	var testInsertBuffer = []byte(`{
		"productid":"123", 
		"productprice": "10.99",
		"productcurrency": "USD",
	}`)

	// Create a new request
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
		t.Errorf("handler return a bad status code: GOT %v WANTED %v",
			status, http.StatusOK)
	}
}

func TestUpdateProductEndpoint(t *testing.T) {

}

func TestProductNameEndpoint(t *testing.T) {

}
