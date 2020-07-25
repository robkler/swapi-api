package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gocql/gocql"
	"log"
	"net/http"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type ErrorJson struct {
	Error string `json:"error"`
}

type GetPlanets struct {
	Planets []Planet `json:"planets" validate:"required"`
}

func validate(s interface{}) error {
	validate := validator.New()
	return validate.Struct(s)
}

func (pr *PlanetRoutes) InsertPlanet(c *gin.Context) {
	defer timeTrack(time.Now(), "Insert planet")
	var planet Planet

	if err := c.BindJSON(&planet); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := validate(planet)

	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = pr.PlanetDb.FindByName(planet.Name)
	if err == nil {
		c.AbortWithStatus(http.StatusConflict)
		return
	}
	if  err.Error() != "not found"{
		c.AbortWithStatus(http.StatusFailedDependency)
		return
	}
	ok := pr.Swapi.ContainPlanet(planet.Name)

	if !ok {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Non-existent planet"})
		return
	}
	err = pr.PlanetDb.Insert(&planet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}

func (pr *PlanetRoutes) GetPlanets(c *gin.Context) {
	defer timeTrack(time.Now(), "get Planet")
	var err error
	planetList := pr.PlanetDb.SelectAllPlanets()

	var newPlanetList []Planet
	for _, ele := range planetList {
		newEle := ele
		newEle.FilmsAppears, err = pr.Swapi.NumOfAppearances(ele.Name)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
			return
		}
		newPlanetList = append(newPlanetList, newEle)
	}
	planets := GetPlanets{newPlanetList}
	c.JSON(http.StatusOK, planets)
}

func (pr *PlanetRoutes) GetByName(c *gin.Context) {
	defer timeTrack(time.Now(), "get planet by name")

	var err error
	name := c.Param("user_name")
	planet, err := pr.PlanetDb.FindByName(name)
	if err != nil {
		if err.Error() == "not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}

	planet.FilmsAppears, err = pr.Swapi.NumOfAppearances(planet.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, planet)
}

func (pr *PlanetRoutes) GetById(c *gin.Context) {
	defer timeTrack(time.Now(), "get planet by id")

	var err error
	uuid, err := gocql.ParseUUID(c.Param("user_uuid"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
		return
	}
	planet, err := pr.PlanetDb.FindById(uuid)
	if err != nil {
		if err.Error() == "not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return

	}
	planet.FilmsAppears, err = pr.Swapi.NumOfAppearances(planet.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, planet)
}

func (pr *PlanetRoutes) DeletePlanet(c *gin.Context) {
	defer timeTrack(time.Now(), "Delete planet")

	var err error
	uuid, err := gocql.ParseUUID(c.Param("user_uuid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
		return
	}
	planet, err := pr.PlanetDb.FindById(uuid)
	if err != nil {
		if err.Error() == "not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}
	err = pr.PlanetDb.DeletePlanet(&planet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusFailedDependency, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, planet)
}
