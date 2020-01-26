package main

import "github.com/gocql/gocql"

type Planet struct {
	Id      gocql.UUID `json:"id"`
	Name    string     `json:"name"`
	Climate string     `json:"climate"`
	Terrain string     `json:"terrain"`
}

func (p *Planet) Insert() error {
	id := gocql.TimeUUID()
	if err := session.Query(`INSERT INTO swapi.planet (id, name, climate, terrain) VALUES (? ,? ,? ,? )`,
		id, p.Name, p.Climate, p.Terrain).Consistency(
		gocql.One).Exec(); err != nil {
		return err
	}
	p.Id = id
	return nil
}

func (p *Planet) FindById() error {
	if err := session.Query(`SELECT name,climate,terrain FROM swapi.planet WHERE id = ?`,
		p.Id.String()).Consistency(
		gocql.One).Scan(&p.Name, &p.Climate, &p.Terrain); err != nil {
		return err
	}
	return nil
}

func (p *Planet) FindByName() error {
	if err := session.Query(`SELECT id, climate, terrain FROM swapi.planet_by_name WHERE name = ?`,
		p.Name).Consistency(
		gocql.One).Scan(&p.Id, &p.Climate, &p.Terrain); err != nil {
		return err
	}
	return nil
}

func (p *Planet) SelectAllPlanets() []Planet {
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

func (p *Planet) DeletePlanet() error {
	return session.Query(`DELETE FROM swapi.planet WHERE id = ?`, p.Id).Consistency(gocql.One).Exec()
}
