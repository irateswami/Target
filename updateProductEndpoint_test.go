package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
