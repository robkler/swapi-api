package routes

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

func (pr *PlanetRoutes) InsertPlanet(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "Insert planet")
	var p Planet
	var err error
	err = json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		log.Println(err)
		return
	}
	err = validate(p)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: err.Error(),
		})
		return
	}

	err = pr.PlanetDb.FindByName(&p)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		return
	}
	ok := pr.Swapi.ContainPlanet(p.Name)
	if !ok {
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(ErrorJson{
			Error: "Non-existent p",
		}) //todo fix
		return
	}
	err = pr.PlanetDb.Insert(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (pr *PlanetRoutes) GetPlanets(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "get Planet")

	planetList := pr.PlanetDb.SelectAllPlanets()
	for _, ele := range planetList {
		ele.FilmsAppears, _ = pr.Swapi.NumOfAppearances(ele.Name) //todo fix
	}
	json.NewEncoder(w).Encode(planetList)
}

func (pr *PlanetRoutes) GetByName(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "get planet by name")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	p.Name = vars["user_name"]
	err = pr.PlanetDb.FindByName(&p)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	p.FilmsAppears, _ = pr.Swapi.NumOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func (pr *PlanetRoutes) GetById(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "get planet by id")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	p.Id = uuid
	err = pr.PlanetDb.FindById(&p)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	p.FilmsAppears, _ = pr.Swapi.NumOfAppearances(p.Name)
	json.NewEncoder(w).Encode(p)
}

func (pr *PlanetRoutes) DeletePlanet(w http.ResponseWriter, r *http.Request) {
	defer timeTrack(time.Now(), "Delete planet")

	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		return
	}
	p.Id = uuid
	err = pr.PlanetDb.FindById(&p)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
	err = pr.PlanetDb.DeletePlanet(&p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(p)
}
