// Code generated by MockGen. DO NOT EDIT.
// Source: planet.go

// Package mock is a generated GoMock package.
package mock

import (
	gocql "github.com/gocql/gocql"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	routes "swapi/routes"
)

// MockPlanetDb is a mock of PlanetDb interface
type MockPlanetDb struct {
	ctrl     *gomock.Controller
	recorder *MockPlanetDbMockRecorder
}

// MockPlanetDbMockRecorder is the mock recorder for MockPlanetDb
type MockPlanetDbMockRecorder struct {
	mock *MockPlanetDb
}

// NewMockPlanetDb creates a new mock instance
func NewMockPlanetDb(ctrl *gomock.Controller) *MockPlanetDb {
	mock := &MockPlanetDb{ctrl: ctrl}
	mock.recorder = &MockPlanetDbMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPlanetDb) EXPECT() *MockPlanetDbMockRecorder {
	return m.recorder
}

// Insert mocks base method
func (m *MockPlanetDb) Insert(p *routes.Planet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockPlanetDbMockRecorder) Insert(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockPlanetDb)(nil).Insert), p)
}

// FindById mocks base method
func (m *MockPlanetDb) FindById(id gocql.UUID) (routes.Planet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", id)
	ret0, _ := ret[0].(routes.Planet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById
func (mr *MockPlanetDbMockRecorder) FindById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockPlanetDb)(nil).FindById), id)
}

// FindByName mocks base method
func (m *MockPlanetDb) FindByName(name string) (routes.Planet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", name)
	ret0, _ := ret[0].(routes.Planet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByName indicates an expected call of FindByName
func (mr *MockPlanetDbMockRecorder) FindByName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockPlanetDb)(nil).FindByName), name)
}

// SelectAllPlanets mocks base method
func (m *MockPlanetDb) SelectAllPlanets(state []byte) ([]routes.Planet, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAllPlanets", state)
	ret0, _ := ret[0].([]routes.Planet)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// SelectAllPlanets indicates an expected call of SelectAllPlanets
func (mr *MockPlanetDbMockRecorder) SelectAllPlanets(state interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAllPlanets", reflect.TypeOf((*MockPlanetDb)(nil).SelectAllPlanets), state)
}

// DeletePlanet mocks base method
func (m *MockPlanetDb) DeletePlanet(p *routes.Planet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePlanet", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePlanet indicates an expected call of DeletePlanet
func (mr *MockPlanetDbMockRecorder) DeletePlanet(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePlanet", reflect.TypeOf((*MockPlanetDb)(nil).DeletePlanet), p)
}

// MockSwapi is a mock of Swapi interface
type MockSwapi struct {
	ctrl     *gomock.Controller
	recorder *MockSwapiMockRecorder
}

// MockSwapiMockRecorder is the mock recorder for MockSwapi
type MockSwapiMockRecorder struct {
	mock *MockSwapi
}

// NewMockSwapi creates a new mock instance
func NewMockSwapi(ctrl *gomock.Controller) *MockSwapi {
	mock := &MockSwapi{ctrl: ctrl}
	mock.recorder = &MockSwapiMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSwapi) EXPECT() *MockSwapiMockRecorder {
	return m.recorder
}

// NumOfAppearances mocks base method
func (m *MockSwapi) NumOfAppearances(planet string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NumOfAppearances", planet)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NumOfAppearances indicates an expected call of NumOfAppearances
func (mr *MockSwapiMockRecorder) NumOfAppearances(planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NumOfAppearances", reflect.TypeOf((*MockSwapi)(nil).NumOfAppearances), planet)
}

// ContainPlanet mocks base method
func (m *MockSwapi) ContainPlanet(planet string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainPlanet", planet)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ContainPlanet indicates an expected call of ContainPlanet
func (mr *MockSwapiMockRecorder) ContainPlanet(planet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainPlanet", reflect.TypeOf((*MockSwapi)(nil).ContainPlanet), planet)
}
