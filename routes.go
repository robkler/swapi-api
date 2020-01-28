package main

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type ErrorJson struct {
	Error string `json:"error"`
}

func validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Insert planet")
	var planet Planet
	var err error
	err = json.NewDecoder(r.Body).Decode(&planet)

	if err != nil {
		w.WriteHeader(BadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		log.Println(err)
		return
	}
	err = validate(planet)

	if err != nil {
		log.Println(err)
		w.WriteHeader(BadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		return
	}

	err = planet.FindByName()
	if err == nil {
		w.WriteHeader(Conflict)
		return
	}
	ok := planets.containPlanet(planet.Name)
	if !ok {
		w.WriteHeader(PreconditionFailed)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: "Non-existent planet",
		}) //todo fix
		return
	}
	err = planet.Insert()
	if err != nil {
		w.WriteHeader(InternalServerError)
		return
	}
	w.WriteHeader(Insert)
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get Planets")

	var planet Planet
	planetList := planet.SelectAllPlanets()
	for _, ele := range planetList {
		ele.FilmsAppears, _ = planets.numOfAppearances(ele.Name) //todo fix
	}
	json.NewEncoder(w).Encode(planetList)
}

func GetByName(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get planet by name")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	p.Name = vars["user_name"]
	err = p.FindByName()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(NotFound)
			return
		}
	}
	p.FilmsAppears, _ = planets.numOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get planet by id")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(PreconditionFailed)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(NotFound)
			return
		}
	}
	p.FilmsAppears, _ = planets.numOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Delete planet")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(PreconditionFailed)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(NotFound)
			return
		}
	}
	err = p.DeletePlanet()
	if err != nil {
			w.WriteHeader(PreconditionFailed)
			return
	}
	json.NewEncoder(w).Encode(p)
}

const (
	Insert int = 201
	BadRequest int = 400
	NotFound int = 404
	Conflict int = 409
	PreconditionFailed int = 412
	InternalServerError int = 500
)