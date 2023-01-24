package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"task_api/internal/task/handlers"
	"task_api/internal/task/services"
)

func main() {
	apiService := services.NewApiService()
	apiHandler := handlers.NewApiHandler(apiService)

	log.Println("Listen on port: 8080")

	r := mux.NewRouter()
	r.HandleFunc("/", apiHandler.Task).Methods(http.MethodPost)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Route not found"))
		return
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
