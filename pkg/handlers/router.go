package handlers

import "github.com/gorilla/mux"

func NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/",Home).Methods("GET")
	r.HandleFunc("/item/{id}",ShowItem).Methods("GET")

	return r
}
