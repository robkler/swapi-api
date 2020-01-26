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
	router.HandleFunc("/", GetPlants).Methods("GET")
	router.HandleFunc("/", GetByName).Methods("GET")
	router.HandleFunc("/id/{user_uuid}", GetById).Methods("GET")
	router.HandleFunc("/", DeletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}


func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	json.NewDecoder(r.Body).Decode(&planet)
	planet.Id = planet.Insert(session)
	json.NewEncoder(w).Encode(planet)
}


func GetPlants(w http.ResponseWriter, r *http.Request) {

}
func GetByName(w http.ResponseWriter, r *http.Request) {
}
func GetById(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil{
		w.WriteHeader(412)
		return
	}
	p.Id = uuid
	fmt.Println("user_id",p.Id)
	err = p.FindById(session)
	if err != nil {
		if err.Error() == "not found"{
			w.WriteHeader(404)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}
func DeletePlanet(w http.ResponseWriter, r *http.Request) {
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

func (p *Planet) FindById(session *gocql.Session) error {
	if err := session.Query(`SELECT name,climate,terrain FROM swapi.planeta WHERE id = ?`,
		p.Id.String()).Consistency(
		gocql.One).Scan(&p.Name,&p.Climate,&p.Terrain); err != nil {
		return err
	}
	fmt.Println("user_id",gocql.TimeUUID())
		return nil
}
