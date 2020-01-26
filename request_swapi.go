package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Res struct {
	Results []Planets `json:"results"`
	Next    string    `json:"next"`
}
type Planets struct {
	Name    string   `json:"name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"residents"`
}

type MapPlanets struct {
	planets map[string]Planets
}

var planets MapPlanets

func getAllPlanets() {
	log.Println("Getting Planets")
	mapPlanets := make(map[string]Planets)
	planets = MapPlanets{
		planets: mapPlanets,
	}
	get("https://swapi.co/api/planets/")
	log.Println("Got Planets")
}

func get(next string) {
	var r = Res{}
	var err error
	rep, err := http.Get(next)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(rep.Body).Decode(&r)
	for _, ele := range r.Results {
		planets.planets[ele.Name] = ele
	}
	if r.Next != "" {
		get(r.Next)
	}
}

func (m *MapPlanets) numOfAppearances(planet string) (int, error) {
	if !m.containPlanet(planet) {
		return 0, errors.New("Non-existent planet")
	}
	return len(m.planets[planet].Films), nil
}

func (m *MapPlanets) containPlanet(planet string) bool {
	_, contain := m.planets[planet]
	return contain
}
