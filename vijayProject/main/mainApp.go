package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vijay746/master_fbApp/vijayProject/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterEventRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
