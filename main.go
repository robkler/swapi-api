package main

import (
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster(cassandraHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cassandraUserName,
		Password: cassandraPassword,
	}
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	getAllPlanets()
	log.Println("Api Ready")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/planet", InsertPlanet).Methods("POST")
	router.HandleFunc("/planet", GetPlanets).Methods("GET")
	router.HandleFunc("/planet/{user_name}/name", GetByName).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}/id", GetById).Methods("GET")
	router.HandleFunc("/planet/{user_uuid}", DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+apiPort, router))
}
