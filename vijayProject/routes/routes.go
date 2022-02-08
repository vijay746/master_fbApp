package routes

import (
	"github.com/gorilla/mux"
	"github.com/vijay746/master_fbApp/vijayProject/controllers"
)

var RegisterEventRoutes = func(router *mux.Router) {
	router.HandleFunc("/createPost/", controllers.CreatePost).Methods("POST")
	router.HandleFunc("/getAllPosts/", controllers.GetPost).Methods("GET")
	router.HandleFunc("/getPostById/{}", controllers.GetPostById).Methods("GET")

	router.HandleFunc("/createComment/", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/getAllComments/", controllers.GetComment).Methods("GET")

	router.HandleFunc("/createLike/", controllers.CreateLike).Methods("PUT")
	router.HandleFunc("/getAllLikeCount/", controllers.GetLike).Methods("GET")

	router.HandleFunc("/createHeart/", controllers.CreateHeart).Methods("PUT")
	router.HandleFunc("/getAllHeartCount/", controllers.GetHeart).Methods("GET")
}
