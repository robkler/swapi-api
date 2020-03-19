package routes

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"swapi/db"
	"swapi/swapi"
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
	var planet db.Planet
	var err error
	err = json.NewDecoder(r.Body).Decode(&planet)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		log.Println(err)
		return
	}
	err = validate(planet)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		return
	}

	err = planet.FindByName()
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	ok := request_swapi.Planets.ContainPlanet(planet.Name)
	if !ok {
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: "Non-existent planet",
		}) //todo fix
		return
	}
	err = planet.Insert()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get Planet")

	var planet db.Planet
	planetList := planet.SelectAllPlanets()
	for _, ele := range planetList {
		ele.FilmsAppears, _ = request_swapi.Planets.NumOfAppearances(ele.Name) //todo fix
	}
	json.NewEncoder(w).Encode(planetList)
}

func GetByName(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get planet by name")

	p := db.Planet{}
	var err error
	vars := mux.Vars(r)
	p.Name = vars["user_name"]
	err = p.FindByName()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	p.FilmsAppears, _ = request_swapi.Planets.NumOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Get planet by id")

	p := db.Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	p.FilmsAppears, _ = request_swapi.Planets.NumOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(),"Delete planet")

	p := db.Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	p.Id = uuid
	err = p.FindById()
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	err = p.DeletePlanet()
	if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
	}
	json.NewEncoder(w).Encode(p)
}

