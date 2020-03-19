package db

import (
	"github.com/gocql/gocql"
	"log"
	"swapi/environment"
)

var session *gocql.Session

func init() {
	var err error
	cluster := gocql.NewCluster(environment.CassandraHost())
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: environment.CassandraUserName(),
		Password: environment.CassandraPassword(),
	}
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("PlanetDb Connected")
}
