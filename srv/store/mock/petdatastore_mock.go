// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jianhan/petstore_ms/srv/store/datastore (interfaces: PetDataStore)

// Package mock is a generated GoMock package.
package mock

import (
	gomock "github.com/golang/mock/gomock"
	pet "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	reflect "reflect"
)

// MockPetDataStore is a mock of PetDataStore interface
type MockPetDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockPetDataStoreMockRecorder
}

// MockPetDataStoreMockRecorder is the mock recorder for MockPetDataStore
type MockPetDataStoreMockRecorder struct {
	mock *MockPetDataStore
}

// NewMockPetDataStore creates a new mock instance
func NewMockPetDataStore(ctrl *gomock.Controller) *MockPetDataStore {
	mock := &MockPetDataStore{ctrl: ctrl}
	mock.recorder = &MockPetDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPetDataStore) EXPECT() *MockPetDataStoreMockRecorder {
	return m.recorder
}

// FindPetById mocks base method
func (m *MockPetDataStore) FindPetById(arg0 int64) (*pet.Pet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPetById", arg0)
	ret0, _ := ret[0].(*pet.Pet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPetById indicates an expected call of FindPetById
func (mr *MockPetDataStoreMockRecorder) FindPetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPetById", reflect.TypeOf((*MockPetDataStore)(nil).FindPetById), arg0)
}

// InsertPet mocks base method
func (m *MockPetDataStore) InsertPet(arg0 *pet.Pet) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPet", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertPet indicates an expected call of InsertPet
func (mr *MockPetDataStoreMockRecorder) InsertPet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPet", reflect.TypeOf((*MockPetDataStore)(nil).InsertPet), arg0)
}

// UpdatePet mocks base method
func (m *MockPetDataStore) UpdatePet(arg0 *pet.UpdatePetRequest) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePet", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePet indicates an expected call of UpdatePet
func (mr *MockPetDataStoreMockRecorder) UpdatePet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePet", reflect.TypeOf((*MockPetDataStore)(nil).UpdatePet), arg0)
}
