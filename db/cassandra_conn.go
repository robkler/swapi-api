package db

import (
	"github.com/gocql/gocql"
	"log"
)


func (db *PlanetDb) Init() {
	var err error
	cluster := gocql.NewCluster(db.Config.CassandraHost)
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: db.Config.CassandraUsername,
		Password: db.Config.CassandraPassword,
	}
	db.session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PlanetDb Connected")
}
