package test

import (
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"testing"
)

func TestGetPlanetByName(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}

	req, err := http.NewRequest("GET", "planet/Name",nil)
	if err != nil {
		t.Fatal(err)
	}
	q :=req.URL.Query()
	q.Add("user_name", "Name")
	req.URL.RawQuery = q.Encode()
	rec := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.GetByName)
	handler.ServeHTTP(rec, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}