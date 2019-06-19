package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jianhan/petstore_ms/srv/store/mock"
	store "github.com/jianhan/petstore_ms/srv/store/proto/pet"
)

func TestInsertPet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	insertRequest := store.InsertPetRequest{
		Pet: &store.Pet{
			Name:      "test pet",
			PhotoUrls: []string{"url1"},
			Status:    store.Pending.String(),
		},
	}
	petService := mock.NewMockPetService(ctrl)
	petService.EXPECT().InsertPet(context.Background(), &insertRequest).Return(&store.InsertPetResponse{
		Pet: &store.Pet{
			Id:        1,
			Name:      "test pet",
			PhotoUrls: []string{"url1"},
			Status:    store.Pending.String(),
		},
	}, nil)

	petHandler := &petHandler{petService: petService}

	req, err := http.NewRequest(http.MethodPost, "/api/v2/pet", strings.NewReader(`{"pet":{"name":"test pet","photo_urls":["url1"],"status":"pending"}}`))
	if err != nil {
		t.Fatalf("not expected err, got '%s'", err)
	}

	rr := httptest.NewRecorder()

	http.HandlerFunc(petHandler.InsertPet).ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Fatalf("invalid status code, expected %d, got %d", http.StatusOK, status)
	}

	expectedBody := `{"id":1,"name":"test pet","status":"pending","photo_urls":["url1"]}`
	if actualBody := rr.Body.String(); actualBody != expectedBody {
		t.Fatalf("invalid response content, expected %s, got %s", expectedBody, actualBody)
	}
}
