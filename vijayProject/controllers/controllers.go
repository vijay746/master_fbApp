package controllers

import "net/http"

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for getiing all posts."))
}

func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for getiing a post by its ID."))
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for creating a post."))
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for creating a comment."))
}

func GetComment(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for getting all comments to a post."))
}

func CreateLike(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for like to a post."))
}

func GetLike(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for geting all like count to a post."))
}

func CreateHeart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for heart a post."))
}

func GetHeart(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi this is for geting all heart count to a post."))
}
