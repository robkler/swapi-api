package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"swapi/db"
	"swapi/environment"
	"swapi/routes"
	request_swapi "swapi/swapi"
)


func init() {
	log.Println("Api Ready")
}

func main() {
	planetDb := db.PlanetDb{}
	mapPlanets := request_swapi.MapPlanets{}
	mapPlanets.GetAllPlanets()
	planetRoutes := routes.PlanetRoutes{
		PlanetDb:&planetDb,
		Swapi: &mapPlanets,
	}

	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/planet", planetRoutes.InsertPlanet).Methods("POST")
	router.HandleFunc("/planet", planetRoutes.GetPlanets).Methods("GET")
	router.HandleFunc("/planet/{user_name}/name", planetRoutes.GetByName).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}/id", planetRoutes.GetById).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}", planetRoutes.DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+environment.ApiPort(), router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
