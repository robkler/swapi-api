package test

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"swapi/routes/mock"
	"testing"
)

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	uuid, err := gocql.ParseUUID("572e9d29-a08e-43db-bb80-d393f769de6d")

	findById := db.EXPECT().FindById(uuid).Return(p, nil)
	db.EXPECT().DeletePlanet(gomock.Any()).Return(nil).After(findById)

	req, err := http.NewRequest("DELETE", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.DELETE("/:user_uuid", pr.DeletePlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNoContent)
	}
}

func TestDeleteNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	uuid, err := gocql.ParseUUID("572e9d29-a08e-43db-bb80-d393f769de6d")

	db.EXPECT().FindById(uuid).Return(p, errors.New("not found"))

	req, err := http.NewRequest("DELETE", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.DELETE("/:user_uuid", pr.DeletePlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestDeleteStatusFailed(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	uuid, err := gocql.ParseUUID("572e9d29-a08e-43db-bb80-d393f769de6d")

	db.EXPECT().FindById(uuid).Return(p, errors.New(""))

	req, err := http.NewRequest("DELETE", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.DELETE("/:user_uuid", pr.DeletePlanet)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}