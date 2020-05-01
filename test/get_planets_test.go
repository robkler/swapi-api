package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"testing"
)

func TestGetPlanets(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}

	var jsonStr = []byte(``)
	req, err := http.NewRequest("GET", "",bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.GetPlanets)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}