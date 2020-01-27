package main

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type ErrorJson struct {
	Error string `json:"error"`
}

func validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	defer log.Println("Insert Planet took: ", timeDuration(t1))
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
		err = errors.New("Non-existent planet")
		json.NewEncoder(w).Encode(err) //todo fix
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
			w.WriteHeader(NotFound)
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
	t1 := time.Now()
	defer log.Println("Delete Planet took: ", timeDuration(t1))
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

func timeDuration(t time.Time) time.Duration {
	t2 := time.Now()
	return t2.Sub(t)
}

const (
	Insert int = 201
	BadRequest int = 400
	NotFound int = 404
	Conflict int = 409
	PreconditionFailed int = 412
	InternalServerError int = 500
)