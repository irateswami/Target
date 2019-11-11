package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func test_main(t *testing.T) {

	// Create a request to pass to the handler
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)

	// Serve it up
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code, got %v WANTED: %v", status, http.StatusOK)
	}

	// Check to make sure the response was what we wanted
	expected := "welcome home!"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v WANTED: %v", rr.Body.String(), expected)
	}

}
