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

func TestGetPlanetById(t *testing.T) {
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
	s.EXPECT().NumOfAppearances(gomock.Any()).Return(1,nil).After(findById)

	req, err := http.NewRequest("GET", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_uuid", pr.GetById)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetPlanetByIdNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	db.EXPECT().FindById(gomock.Any()).Return(p, errors.New("not found"))
	req, err := http.NewRequest("GET", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_uuid", pr.GetById)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}


func TestGetPlanetByIdDataBaseError(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	findById := db.EXPECT().FindById(gomock.Any()).Return(p, nil)
	s.EXPECT().NumOfAppearances(gomock.Any()).Return(1,errors.New("")).After(findById)
	req, err := http.NewRequest("GET", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_uuid", pr.GetById)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}

func TestGetPlanetByIdSwapiErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	db := mock.NewMockPlanetDbClient(ctrl)
	s := mock.NewMockSwapiClient(ctrl)
	pr := routes.PlanetRoutes{
		PlanetDb: db,
		Swapi:    s,
	}
	p := routes.Planet{}
	db.EXPECT().FindById(gomock.Any()).Return(p, errors.New(""))
	req, err := http.NewRequest("GET", "/572e9d29-a08e-43db-bb80-d393f769de6d", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/:user_uuid", pr.GetById)
	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFailedDependency {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFailedDependency)
	}
}
