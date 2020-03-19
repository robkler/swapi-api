package request_swapi

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type Res struct {
	Results []Planet `json:"results"`
	Next    string   `json:"next"`
}
type Planet struct {
	Name    string   `json:"name"`
	Climate string   `json:"climate"`
	Terrain string   `json:"terrain"`
	Films   []string `json:"residents"`
}

type MapPlanets struct {
	planets map[string]Planet
}

var Planets MapPlanets

func GetAllPlanets() {
	log.Println("Getting Planet")
	mapPlanets := make(map[string]Planet)
	Planets = MapPlanets{
		planets: mapPlanets,
	}
	Get("https://swapi.co/api/planets/")
	log.Println(Planets)
	log.Println("Got Planet")
}

func Get(next string) {
	var r = Res{}
	var err error
	rep, err := http.Get(next)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(rep.Body).Decode(&r)
	for _, ele := range r.Results {
		Planets.planets[ele.Name] = ele
	}
	if r.Next != "" {
		Get(r.Next)
	}
}

func (m *MapPlanets) NumOfAppearances(planet string) (int, error) {
	if !m.ContainPlanet(planet) {
		return 0, errors.New("Non-existent planet")
	}
	return len(m.planets[planet].Films), nil
}

func (m *MapPlanets) ContainPlanet(planet string) bool {
	_, contain := m.planets[planet]
	return contain
}
