package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
