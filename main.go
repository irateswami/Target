package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
)

// Because we have one example and I'm not totally sure on how the api url string is constructed, this will just be a constant. In the real world this would be constructed via arguments very closely to how the construct_product url is constructed and handled.
const redsky_url string = "https://redsky.target.com/v2/pdp/tcin/13860428?excludes=taxonomy,price,promotion,bulk_ship,rating_and_review_reviews,rating_and_review_statistics,question_answer_statistics"

type Product struct {
	id       string
	name     string
	price    string
	currency string
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome home!")
}

func (p *Product) product_name_api(u string) {

	// Construct http request to the product info api. In this case the api address is a string constant, but in the real world this would be construct via another function
	response, err := http.Get(u)
	if err != nil {
		log.Fatal("product information api request failed: ", err)
	}
	defer response.Body.Close()

	// Read in the data using io
	unstring_data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("ioutil readall failed: ", err)
	}

	// Stringify the data
	data := string(unstring_data)

	// Big thanks to Josh Baker for making this json parser with syntax that actually makes sense
	product_name := gjson.Get(data, "product.item.product_description.title").String()

	p.name = product_name

}

func construct_product(w http.ResponseWriter, r *http.Request) {

	// Grab all the variables passed
	vars := mux.Vars(r)

	// Let everyone know things went okay
	w.WriteHeader(http.StatusOK)

	//fmt.Fprintf(w, "Product: %v\n", vars["id"])
	// Allocate memory for a new product
	product := new(Product)

	product.id = vars["id"]
	product.product_name_api(redsky_url)
	fmt.Fprintf(w, "", product.name)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	// If just a slash, be welcoming
	router.HandleFunc("/", home)

	// If searching for product by id, be informative
	router.HandleFunc("/product/{id:[0-9]+}", construct_product)

	// Serve up everything on localhost:8080
	log.Fatal(http.ListenAndServe(":8080", router))

}
