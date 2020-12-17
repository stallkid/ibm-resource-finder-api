package controller

import (
	"encoding/json"
	"ibm-resource-finder-api/service"
	"net/http"

	"github.com/gorilla/mux"
)

func GetResourceByW3id(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["email"]

	response := service.GetResourceData(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetResourceByManagerDn(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["dn"]

	response := service.GetManagerResourcesData(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetResourceSkills(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["uid"]

	response := service.GetResourceDetailsByUid(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetResourceByManagerDnWithDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["dn"]

	response := service.GetResourceByManagerDnWithDetails(key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
