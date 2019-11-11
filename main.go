package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":8080", router))

}
