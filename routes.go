package main

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"net/http"
)

func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	var err error
	json.NewDecoder(r.Body).Decode(&planet)
	err = planet.FindByName()
	if err == nil {
		w.WriteHeader(409)
		return
	}
	err = planet.Insert()
	if err != nil {
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(planet)
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	planetList := planet.SelectAllPlanets()
	json.NewEncoder(w).Encode(planetList)

}
func GetByName(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	p.Name = vars["user_name"]
	err = p.FindByName()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}
func GetById(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(412)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}
func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(412)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err.Error() == "not found"{
		w.WriteHeader(404)
		return
	}
	err = p.DeletePlanet()
	if err != nil {
		if err != nil {
			w.WriteHeader(412)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}