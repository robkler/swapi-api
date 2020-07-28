package test

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"swapi/routes/mock"
	"testing"
)

func TestGetPlanetByName(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDb(ctrl)
	s := mock.NewMockPlanetDb(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	findByName := db.EXPECT().FindByName("test").Return(p, nil)
	s.EXPECT().NumOfAppearances(gomock.Any()).Return(1,nil).After(findByName)
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_name", pr.GetByName)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		fmt.Println(rr.Body.String())
		t.Errorf("Response code should be 200, was: %d", rr.Code)
	}
}

func TestGetPlanetByNameNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDb(ctrl)
	s := mock.NewMockPlanetDb(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	db.EXPECT().FindByName("test").Return(p, errors.New("not found"))
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_name", pr.GetByName)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}


func TestGetPlanetByNameDataBaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDb(ctrl)
	s := mock.NewMockPlanetDb(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	 db.EXPECT().FindByName("test").Return(p, errors.New(""))
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_name", pr.GetByName)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}

func TestGetPlanetByNameSwapiErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDb(ctrl)
	s := mock.NewMockPlanetDb(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	findByName := db.EXPECT().FindByName("test").Return(p, nil)
	s.EXPECT().NumOfAppearances(gomock.Any()).Return(1,errors.New("")).After(findByName)
	req, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_name", pr.GetByName)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}
