package main

import (
	"encoding/json"
	"ibm-resource-finder-api/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// função principal
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/resource/{email}", controller.GetResourceByW3id).Methods("GET")
	router.HandleFunc("/manager/resource/{dn}", controller.GetResourceByManagerDn).Methods("GET")
	router.HandleFunc("/resource/details/{uid}", controller.GetResourceDetails).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getResource(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Hello World")
}
