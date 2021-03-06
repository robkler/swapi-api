package db

import (
	"github.com/gocql/gocql"
	"swapi/environment"
	"swapi/routes"
)

type PlanetDb struct{
	Config  environment.CassandraConfig
	session *gocql.Session
}
func (db *PlanetDb) Insert(p *routes.Planet) error {
	id := gocql.TimeUUID()
	if err := db.session.Query(`INSERT INTO swapi.planet (id, name, climate, terrain) VALUES (? ,? ,? ,? )`,
		id, p.Name, p.Climate, p.Terrain).Consistency(
		gocql.Quorum).Exec(); err != nil {
		return err
	}
	p.Id = id
	return nil
}

func (db *PlanetDb) FindById(id gocql.UUID) (routes.Planet, error) {
	p := routes.Planet{Id: id}
	if err := db.session.Query(`SELECT name,climate,terrain FROM swapi.planet WHERE id = ?`,
		p.Id.String()).Consistency(
		gocql.Quorum).Scan(&p.Name, &p.Climate, &p.Terrain); err != nil {
		return p, err
	}
	return p, nil
}

func (db *PlanetDb) FindByName(name string) (routes.Planet, error) {
	p := routes.Planet{Name: name}
	if err := db.session.Query(`SELECT id, climate, terrain FROM swapi.planet_by_name WHERE name = ?`,
		p.Name).Consistency(
		gocql.Quorum).Scan(&p.Id, &p.Climate, &p.Terrain); err != nil {
		return p, err
	}
	return p, nil
}

func (db *PlanetDb) SelectAllPlanets(state []byte) ([]routes.Planet, []byte){
	var planetList []routes.Planet
	m := map[string]interface{}{}
	i := db.session.Query(`SELECT id, name,climate,terrain FROM swapi.planet_by_name`).Consistency(
		gocql.Quorum).PageSize(10).PageState(state).Iter()

	s := i.PageState()
	for i.MapScan(m) {
		planetList = append(planetList, routes.Planet{
			Id:      m["id"].(gocql.UUID),
			Name:    m["name"].(string),
			Climate: m["climate"].(string),
			Terrain: m["terrain"].(string),
		})
		m = map[string]interface{}{}
	}
	return planetList , s
}

func (db *PlanetDb) DeletePlanet(p *routes.Planet) error {
	return db.session.Query(`DELETE FROM swapi.planet WHERE id = ?`, p.Id).Consistency(gocql.Quorum).Exec()
}
