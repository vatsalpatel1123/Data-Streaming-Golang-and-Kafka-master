package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"your_project/handlers"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/stream/start", handlers.StartStream).Methods("POST")
	router.HandleFunc("/stream/{stream_id}/send", handlers.SendStreamData).Methods("POST")
	router.HandleFunc("/stream/{stream_id}/results", handlers.StreamResults).Methods("GET")
	return router
}
