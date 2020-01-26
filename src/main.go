package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type heartbeatResponse struct {
	Status string `json:"status"`
	Code int `json:"code"`

}

func heartbeat(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(heartbeatResponse{Status: "OK", Code: 200})
}


var session *gocql.Session

func main() {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "escrow"
	cluster.Consistency = gocql.One
	session, _ = cluster.CreateSession()

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", heartbeat).Methods("GET")
	router.HandleFunc("/", InsertPlanet).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	json.NewDecoder(r.Body).Decode(&planet)
	planet.Id = planet.Insert(session)
	json.NewEncoder(w).Encode(planet)
}

type Planet struct {
	Id gocql.UUID `json:"id"`
	Name string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

func (p *Planet) Insert(session *gocql.Session) gocql.UUID {
	if err := session.Query(`INSERT INTO swapi.planeta (id, name, climate, terrain) VALUES (? ,? ,? ,? )`,
		gocql.TimeUUID(),p.Name,p.Climate,p.Terrain).Consistency(
		gocql.One).Exec(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("user_id",gocql.TimeUUID())
	return gocql.TimeUUID()
}
