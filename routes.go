package main

import (
	"encoding/json"
	"errors"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	defer log.Println("Insert Planet took: ", timeDuration(t1))
	var planet Planet
	var err error
	err = json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		w.WriteHeader(412) //todo melhorar esse erro
		log.Println(err)
		return
	}
	err = planet.FindByName()
	if err == nil {
		w.WriteHeader(409)
		return
	}
	ok := planets.containPlanet(planet.Name)
	if !ok {
		w.WriteHeader(412)
		err = errors.New("Non-existent planet")
		json.NewEncoder(w).Encode(err) //todo fix
		return
	}
	err = planet.Insert()
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(201)
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	defer log.Println("Get Planets took: ", timeDuration(t1))
	var planet Planet
	planetList := planet.SelectAllPlanets()
	for _, ele := range planetList {
		ele.FilmsAppears, _ = planets.numOfAppearances(ele.Name) //todo fix
	}
	json.NewEncoder(w).Encode(planetList)

}
func GetByName(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	defer log.Println("Get By name took: ", timeDuration(t1))
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
	p.FilmsAppears, _ = planets.numOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}
func GetById(w http.ResponseWriter, r *http.Request) {

	t1 := time.Now()
	defer log.Println("Get by id took: ", timeDuration(t1))
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
	p.FilmsAppears, _ = planets.numOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}
func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	defer log.Println("Delete Planet took: ", timeDuration(t1))
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
	err = p.DeletePlanet()
	if err != nil {
			w.WriteHeader(412)
			return
	}
	json.NewEncoder(w).Encode(p)
}

func timeDuration(t time.Time) time.Duration {
	t2 := time.Now()
	return t2.Sub(t)
}
