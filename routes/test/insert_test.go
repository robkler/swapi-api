package test

import (
	"bytes"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"swapi/routes/mock"
	"testing"
)

func TestSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	findByName := db.EXPECT().FindByName("Name").Return(p, nil)
	contain := s.EXPECT().ContainPlanet("Name").Return(true, nil).After(findByName)
	db.EXPECT().Insert(gomock.Any()).Return(nil).After(contain)
	var jsonStr = []byte(`{
			"name": "Name",
			"climate":"Climate",
			"terrain": "Terrain"
			}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", pr.InsertPlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func TestWrongJson(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	var jsonStr = []byte(`{
			"name": "Name2",
			"climate":"Climate"
			`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", pr.InsertPlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestJsonIncomplete(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	var jsonStr = []byte(`{
			"name": "Name2",
			"climate":"Climate"
			}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", pr.InsertPlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestPlanetDontExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	findByName := db.EXPECT().FindByName("Name").Return(p, nil)
	s.EXPECT().ContainPlanet("Name").Return(false, nil).After(findByName)

	var jsonStr = []byte(`{
			"name": "Name",
			"climate":"Climate",
			"terrain": "Terrain"
			}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", pr.InsertPlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}

func TestPlanetDbErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	db.EXPECT().FindByName("Name").Return(p, errors.New(""))

	var jsonStr = []byte(`{
			"name": "Name",
			"climate":"Climate",
			"terrain": "Terrain"
			}`)

	req, err := http.NewRequest("POST", "/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.POST("/", pr.InsertPlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}