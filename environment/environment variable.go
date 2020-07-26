package environment

import (
	"os"
)

type CassandraConfig struct {
	CassandraHost string
	CassandraUsername string
	CassandraPassword string
}

var cassandraConfig = CassandraConfig{}
func init() {
	cassandraConfig.CassandraHost = os.Getenv("CASSANDRA_HOST")
	cassandraConfig.CassandraUsername = os.Getenv("CASSANDRA_USERNAME")
	cassandraConfig.CassandraPassword = os.Getenv("CASSANDRA_PASSWORD")
}

var apiPort  = os.Getenv("API_PORT")

func ApiPort() string {
	return apiPort
}
func CassandraHost() string {
	return cassandraConfig.CassandraHost
}
func CassandraUserName() string {
	return cassandraConfig.CassandraUsername
}
func CassandraPassword() string {
	return cassandraConfig.CassandraPassword
}