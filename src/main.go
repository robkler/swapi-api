package main

import (
	"encoding/json"
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

func InsertPlanet(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	var err error
	json.NewDecoder(r.Body).Decode(&planet)
	err = planet.FindByName(session)
	if err == nil {
		w.WriteHeader(409)
		return
	}
	err = planet.Insert(session)
	if err != nil {
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(planet)
}

func GetPlanets(w http.ResponseWriter, r *http.Request) {
	var planet Planet
	planetList := planet.SelectAllPlanets(session)
	json.NewEncoder(w).Encode(planetList)

}
func GetByName(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	p.Name = vars["user_name"]
	err = p.FindByName(session)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}
func GetById(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(412)
		return
	}
	p.Id = uuid
	err = p.FindById(session)
	if err != nil {
		if err.Error() == "not found" {
			w.WriteHeader(404)
			return
		}
	}
	json.NewEncoder(w).Encode(p)
}
func DeletePlanet(w http.ResponseWriter, r *http.Request) {
	p := Planet{}
	var err error
	vars := mux.Vars(r)
	uuid, err := gocql.ParseUUID(vars["user_uuid"])
	if err != nil {
		w.WriteHeader(412)
		return
	}
	p.Id = uuid
	err = p.FindById(session)
	if err.Error() == "not found"{
		w.WriteHeader(404)
		return
	}
	err = p.DeletePlanet(session)
	if err != nil {
		if err != nil {
			w.WriteHeader(412)
			return
		}
	}

	json.NewEncoder(w).Encode(p)
}

type Planet struct {
	Id      gocql.UUID `json:"id"`
	Name    string     `json:"name"`
	Climate string     `json:"climate"`
	Terrain string     `json:"terrain"`
}

func (p *Planet) Insert(session *gocql.Session) error {
	id := gocql.TimeUUID()
	if err := session.Query(`INSERT INTO swapi.planet (id, name, climate, terrain) VALUES (? ,? ,? ,? )`,
		id, p.Name, p.Climate, p.Terrain).Consistency(
		gocql.One).Exec(); err != nil {
		return err
	}
	p.Id = id
	return nil
}

func (p *Planet) FindById(session *gocql.Session) error {
	if err := session.Query(`SELECT name,climate,terrain FROM swapi.planet WHERE id = ?`,
		p.Id.String()).Consistency(
		gocql.One).Scan(&p.Name, &p.Climate, &p.Terrain); err != nil {
		return err
	}
	return nil
}

func (p *Planet) FindByName(session *gocql.Session) error {
	if err := session.Query(`SELECT id, climate, terrain FROM swapi.planet_by_name WHERE name = ?`,
		p.Name).Consistency(
		gocql.One).Scan(&p.Id, &p.Climate, &p.Terrain); err != nil {
		return err
	}
	return nil
}

func (p *Planet) SelectAllPlanets(session *gocql.Session) []Planet {
	var planetList []Planet
	m := map[string]interface{}{}
	iterable := session.Query(`SELECT id, name,climate,terrain FROM swapi.planet_by_name`).Consistency(
		gocql.One).Iter()
	for iterable.MapScan(m) {
		planetList = append(planetList, Planet{
			Id:      m["id"].(gocql.UUID),
			Name:    m["name"].(string),
			Climate: m["climate"].(string),
			Terrain: m["terrain"].(string),
		})
		m = map[string]interface{}{}
	}
	return planetList
}

func (p *Planet) DeletePlanet(session *gocql.Session) error {
	return session.Query(`DELETE FROM swapi.planet WHERE id = ?`, p.Id).Consistency(gocql.One).Exec()
}
