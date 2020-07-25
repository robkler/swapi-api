package routes

import "github.com/gocql/gocql"

//go:generate  go run github.com/golang/mock/mockgen  -package mock -destination=./mock/planet_client.go -source=$GOFILE

type Planet struct {
	Id           gocql.UUID `json:"id"`
	Name         string     `json:"name" validate:"required"`
	Climate      string     `json:"climate" validate:"required"`
	Terrain      string     `json:"terrain" validate:"required"`
	FilmsAppears int        `json:"films_appears"`
}

type PlanetRoutes struct {
	PlanetDb PlanetDb
	Swapi    Swapi
}

type PlanetDb interface {
	Insert(p *Planet) error
	FindById(id gocql.UUID) (Planet, error)
	FindByName(name string) (Planet, error)
	SelectAllPlanets() []Planet
	DeletePlanet(p *Planet) error
}

type Swapi interface {
	NumOfAppearances(planet string) (int, error)
	ContainPlanet(planet string) bool
}
