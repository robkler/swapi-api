package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"swapi/routes"
	"testing"
)

func TestSuccess(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}
	var jsonStr = []byte(`{
			"name": "Name2",
			"climate":"Climate",
			"terrain": "Terrain"
			}`)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.InsertPlanet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}

func TestWrongJson(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}
	var jsonStr = []byte(`{
			"name": "Name2",
			"climate":"Climate"
			`)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.InsertPlanet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestJsonIncomplete(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}
	var jsonStr = []byte(`{
			"name": "Name2",
			"climate":"Climate"
			}`)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.InsertPlanet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func TestPlanetDontExist(t *testing.T) {
	db := DbSuccess{}
	s := SwapiTest{}
	pr := routes.PlanetRoutes{
		PlanetDb: &db,
		Swapi:    &s,
	}
	var jsonStr = []byte(`{
			"name": "Name3",
			"climate":"Climate",
			"terrain": "Terrain"
			}`)

	req, err := http.NewRequest("POST", "", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(pr.InsertPlanet)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusPreconditionFailed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusPreconditionFailed)
	}
}