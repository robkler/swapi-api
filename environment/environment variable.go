package environment

import "os"

var apiPort  = os.Getenv("API_PORT")
var cassandraHost  = os.Getenv("CASSANDRA_HOST")
var cassandraUserName = os.Getenv("CASSANDRA_USERNAME")
var cassandraPassword = os.Getenv("CASSANDRA_PASSWORD")

func ApiPort() string {
	return apiPort
}
func CassandraHost() string {
	return cassandraHost
}
func CassandraUserName() string {
	return cassandraUserName
}
func CassandraPassword() string {
	return cassandraPassword
}