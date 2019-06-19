// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jianhan/petstore_ms/srv/store/proto/pet (interfaces: PetServiceHandler)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	pet "github.com/jianhan/petstore_ms/srv/store/proto/pet"
	reflect "reflect"
)

// MockPetServiceHandler is a mock of PetServiceHandler interface
type MockPetServiceHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPetServiceHandlerMockRecorder
}

// MockPetServiceHandlerMockRecorder is the mock recorder for MockPetServiceHandler
type MockPetServiceHandlerMockRecorder struct {
	mock *MockPetServiceHandler
}

// NewMockPetServiceHandler creates a new mock instance
func NewMockPetServiceHandler(ctrl *gomock.Controller) *MockPetServiceHandler {
	mock := &MockPetServiceHandler{ctrl: ctrl}
	mock.recorder = &MockPetServiceHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPetServiceHandler) EXPECT() *MockPetServiceHandlerMockRecorder {
	return m.recorder
}

// InsertPet mocks base method
func (m *MockPetServiceHandler) InsertPet(arg0 context.Context, arg1 *pet.InsertPetRequest, arg2 *pet.InsertPetResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertPet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertPet indicates an expected call of InsertPet
func (mr *MockPetServiceHandlerMockRecorder) InsertPet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertPet", reflect.TypeOf((*MockPetServiceHandler)(nil).InsertPet), arg0, arg1, arg2)
}

// UpdatePet mocks base method
func (m *MockPetServiceHandler) UpdatePet(arg0 context.Context, arg1 *pet.UpdatePetRequest, arg2 *pet.UpdatePetResponse) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePet", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePet indicates an expected call of UpdatePet
func (mr *MockPetServiceHandlerMockRecorder) UpdatePet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePet", reflect.TypeOf((*MockPetServiceHandler)(nil).UpdatePet), arg0, arg1, arg2)
}
