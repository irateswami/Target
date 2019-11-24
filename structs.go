package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Productid       string             `json:"productid,omitempty" bson:"productid,omitempty"`
	Productname     string             `json:"productname,omitempty" bson:"productname,omitempty"`
	Productprice    string             `json:"productprice,omitempty" bson:"productprice,omitempty"`
	ProductCurrency string             `json:"productcurrency,omitempty" bson:"productcurrency,omitempty"`
}

type Configs struct {
	NAMEURLBEGIN string `json:"nameurlbegin"`
	NAMEURLEND   string `json:"nameurlend"`
	DBURL        string `json:"dburl"`
	NAMEURL      string
}
