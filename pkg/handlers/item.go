package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/serhiy-v/test-api/pkg/models"
)

func (h *BaseHandler)CreateItem(w http.ResponseWriter, r *http.Request){
	ctx := context.Background()
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println("Unable to get item from request")
	}
	err = h.itemRepo.CreateItem(ctx, item)
	if err != nil{
		log.Println("Unable to create")
	}
}

func (h *BaseHandler)ShowItem(w http.ResponseWriter, r *http.Request)  {
	ctx := context.Background()
	params := mux.Vars(r)
	id,err  := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Unable to get id from request")
	}
	res, err := h.itemRepo.GetItem(ctx, id)
	json.NewEncoder(w).Encode(res)
}
