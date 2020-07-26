package main

import (
	"log"
	"swapi/server"
)


func init() {
	log.Println("Api Ready")
}

func main() {
	s := server.Server{}
	r := s.New()

	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}
