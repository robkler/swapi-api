package test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"swapi/routes"
	"swapi/routes/mock"
	"testing"
)

func TestGetPlanets(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	uuid,_ := gocql.RandomUUID()
	p := routes.Planet{
		Id: uuid,
		Name:"Name",
		Climate: "Climate",
		Terrain: "Terrain",

	}
	ps := []routes.Planet{p}
	selectAll := db.EXPECT().SelectAllPlanets().Return(ps)
	s.EXPECT().NumOfAppearances(p.Name).Return(1, nil).After(selectAll)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", pr.GetPlanets)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	body,_  :=  ioutil.ReadAll( rr.Body)
	resp := `{"planets":[{"id":"` + uuid.String() + `","name":"Name","climate":"Climate","terrain":"Terrain","films_appears":1}]}`
	if resp != strings.TrimRight(string(body), "\n")  {
		t.Error()
	}
}


func TestGetPlanetsSwapiErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	uuid,_ := gocql.RandomUUID()
	p := routes.Planet{
		Id: uuid,
		Name:"Planet",
	}
	ps := []routes.Planet{p}
	selectAll := db.EXPECT().SelectAllPlanets().Return(ps)
	s.EXPECT().NumOfAppearances(gomock.Any()).Return(1, errors.New("")).After(selectAll)
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/", pr.GetPlanets)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}