package test

import (
	"github.com/gocql/gocql"
	"swapi/db"
	"swapi/environment"
	"swapi/routes"
	"testing"
)

func TestCassandra(t *testing.T) {
	config := environment.CassandraConfig{
		CassandraHost:     "192.13.131.0",
		CassandraUsername: "cassandra",
		CassandraPassword: "cassandra",
	}
	planetDb := db.PlanetDb{Config: config}
	planetDb.Init()
	id, _ := gocql.RandomUUID()
	planet := routes.Planet{
		Id:           id,
		Name:         "NameTest",
		Climate:      "ClimateTest",
		Terrain:      "TerrainTest",
		FilmsAppears: 0,
	}

	err := planetDb.Insert(&planet)
	if err != nil {
		t.Error()
	}
	p, err := planetDb.FindByName("NameTest")
	if err != nil || p != planet {
		t.Error(err)
	}
	p, err = planetDb.FindById(planet.Id)
	if err != nil || p != planet {
		t.Error(err)
	}
	err = planetDb.DeletePlanet(&p)
	if err != nil {
		t.Error(err)
	}
}
