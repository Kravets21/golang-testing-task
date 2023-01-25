package test

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"task_api/internal/task/handlers"
	"task_api/internal/task/services"
	"testing"
)

const (
	taskEndpoint = "/"
)

func TestTaskApiSuccess(t *testing.T) {

	service := services.NewApiService()
	handler := handlers.NewApiHandler(service)

	request := []int64{1, 2, 50, 61, 4}
	requestBody, _ := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPost, taskEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(taskEndpoint, handler.Task)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code, recorder.Code)
}

func TestTaskApiFail(t *testing.T) {

	service := services.NewApiService()
	handler := handlers.NewApiHandler(service)

	request := "test"
	requestBody, _ := json.Marshal(request)

	req, err := http.NewRequest(http.MethodPost, taskEndpoint, bytes.NewReader(requestBody))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc(taskEndpoint, handler.Task)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusBadRequest, recorder.Code, recorder.Code)
}
