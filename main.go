package main

import (
	"embark/api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	host = "10.84.128.51"
	port = "8081"
)

func main() {
	r := mux.NewRouter()
	doAPIRouting(r)

	fmt.Printf("Init API server on Host: %s, Port: %s\n", host, port)
	http.ListenAndServe(host+":"+port, r)
}

func doAPIRouting(r *mux.Router) {
	apiRouter := r.PathPrefix("/api").Subrouter()
	v1Router := apiRouter.PathPrefix("/v1").Subrouter()

	// API v1 routing
	v1Router.HandleFunc("/embarkation_card{format:\\..+$}", api.EmbarkationCard).Methods("POST").Headers("Content-Type", "application/json")
}
