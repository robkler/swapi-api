package main

import (
	"fmt"
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
	fmt.Println("Init")
}

func main() {


	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", InsertPlanet).Methods("POST")
	router.HandleFunc("/", GetPlanets).Methods("GET")
	router.HandleFunc("/name/{user_name}", GetByName).Methods("GET")
	router.HandleFunc("/id/{user_uuid}", GetById).Methods("GET")
	router.HandleFunc("/{user_uuid}", DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+apiPort, router))
}
