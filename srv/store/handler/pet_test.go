package handler

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/jianhan/petstore_ms/srv/store/mock"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

func TestInsertPet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reqPet := &store.Pet{
		Name:      "test pet",
		Status:    store.Available.String(),
		PhotoUrls: []string{"http://url1.com", "http://url2.com"},
	}

	req := &store.InsertPetRequest{
		Pet: reqPet,
	}

	mockPetDatastore := mock.NewMockPetDataStore(ctrl)
	mockPetDatastore.EXPECT().InsertPet(req.Pet).Return(int64(1), nil)
	mockPetDatastore.EXPECT().FindPetById(int64(1)).Return(reqPet, nil)

	petHandler := NewPetServiceHandler(mockPetDatastore)

	if err := petHandler.InsertPet(context.Background(), req, &store.InsertPetResponse{}); err != nil {
		t.Fatalf("not expected erro. got '%s'", err)
	}
}

func TestInsertPetInvalidRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	req := &store.InsertPetRequest{}
	petHandler := NewPetServiceHandler(mock.NewMockPetDataStore(ctrl))
	if err := petHandler.InsertPet(context.Background(), req, &store.InsertPetResponse{}); err == nil {
		t.Fatal("expected err, got nil")
	}
}
