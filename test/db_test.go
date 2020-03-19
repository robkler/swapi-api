package test

import (
	"swapi/routes"
)

type DbSuccess struct {}

func (db * DbSuccess) Insert(p *routes.Planet) error {
	return nil
}
func (db * DbSuccess) FindById(p *routes.Planet) error {
	p.Name = "Name"
	p.Climate = "Climate"
	p.Terrain = "Terrain"
	return nil
}

func (db * DbSuccess) FindByName(p *routes.Planet) error {
	if p.Name == "Name" {
		p.Climate = "Climate"
		p.Terrain = "Terrain"
		return nil
	}
	return &Error{}
}

func (db * DbSuccess) SelectAllPlanets() []routes.Planet {
	var planetList []routes.Planet
	planetList = append(planetList, routes.Planet{
		Name:    "Name",
		Climate: "Climate",
		Terrain: "Terrain",
	})
	return planetList
}

func (db * DbSuccess) DeletePlanet(p *routes.Planet) error {
	return nil
}