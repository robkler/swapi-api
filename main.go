package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"swapi/db"
	"swapi/routes"
	request_swapi "swapi/swapi"
)


func init() {
	log.Println("Api Ready")
}

func main() {
	planetDb := db.PlanetDb{}
	mapPlanets := request_swapi.MapPlanets{}
	mapPlanets.GetAllPlanets()
	planetRoutes := routes.PlanetRoutes{
		PlanetDb:&planetDb,
		Swapi: &mapPlanets,
	}
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.POST("/planet", planetRoutes.InsertPlanet)
	r.GET("/planet", planetRoutes.GetPlanets)
	r.GET("/planet/{user_name}/name", planetRoutes.GetByName)
	r.GET("/planet/{user_uuid}/id", planetRoutes.GetById)
	r.DELETE("/planet/{user_uuid}", planetRoutes.DeletePlanet)
	err := r.Run("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
}

