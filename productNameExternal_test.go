package main

import (
	"testing"
)

func ProductNameExternal(t *testing.T) {

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
