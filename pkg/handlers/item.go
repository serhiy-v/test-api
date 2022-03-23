package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	db2 "github.com/serhiy-v/test-api/pkg/db"
	"github.com/serhiy-v/test-api/pkg/models"
)

var ctx context.Context

func CreateItem(w http.ResponseWriter, r *http.Request){
	var item models.Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Println("Unable to get item from request")
	}
	err = db2.Database.CreateItem(ctx,item)
	if err != nil{
		log.Println("Unable to create")
	}
}

func ShowItem(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)
	id,err  := strconv.Atoi(params["id"])
	if err != nil {
		log.Println("Unable to get id from request")
	}
	res, err := db2.Database.GetItem(ctx,id)
	json.NewEncoder(w).Encode(res)
}
