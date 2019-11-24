package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

func (p *Product) ProductNameExternal(id string) {
	fmt.Println("ProductName called")

	// Construct the url to grab the name of the product
	configs.NAMEURL = configs.NAMEURLBEGIN + id + configs.NAMEURLEND

	// Construct http request to the product name
	response, err := http.Get(configs.NAMEURL)
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
