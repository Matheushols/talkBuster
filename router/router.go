package router

import (
	"github.com/gorilla/mux"
	"talkBuster/handlers"
)

func Start() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/story", handlers.CreateStoryHandler).Methods("POST")

	return r
}
