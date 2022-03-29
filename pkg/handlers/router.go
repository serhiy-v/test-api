package handlers

import (
	"github.com/gorilla/mux"
	"github.com/serhiy-v/test-api/pkg/models"
)

type BaseHandler struct {
	itemRepo models.ItemRepository
}

func NewBaseHandler(repository models.ItemRepository) *BaseHandler {
	return &BaseHandler{
		itemRepo: repository,
	}

}

func (h *BaseHandler) NewRouter() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/item/{id}",h.ShowItem).Methods("GET")
	r.HandleFunc("/item",h.CreateItem).Methods("POST")

	return r
}
