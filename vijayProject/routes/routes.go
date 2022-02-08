package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {

}

var RegisterEventRoutes = func(router *mux.Router) {
	router.HandleFunc("/comment/", CreateComment).Methods("POST")
}
