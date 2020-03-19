package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"swapi/environment"
	"swapi/routes"
	"swapi/swapi"
)


func init() {
	request_swapi.GetAllPlanets()
	log.Println("Api Ready")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.Use(commonMiddleware)
	router.HandleFunc("/planet", routes.InsertPlanet).Methods("POST")
	router.HandleFunc("/planet", routes.GetPlanets).Methods("GET")
	router.HandleFunc("/planet/{user_name}/name", routes.GetByName).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}/id", routes.GetById).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}", routes.DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+environment.ApiPort(), router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
