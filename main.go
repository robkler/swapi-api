package main

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var session *gocql.Session

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "escrow"
	cluster.Consistency = gocql.LocalQuorum
	session, _ = cluster.CreateSession()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", InsertPlanet).Methods("POST")
	router.HandleFunc("/", GetPlanets).Methods("GET")
	router.HandleFunc("/name/{user_name}", GetByName).Methods("GET")
	router.HandleFunc("/id/{user_uuid}", GetById).Methods("GET")
	router.HandleFunc("/{user_uuid}", DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}