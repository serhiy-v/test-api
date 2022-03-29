package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serhiy-v/test-api/pkg/models"
)

func Home(w http.ResponseWriter, r *http.Request)  {
	item := models.Item{
		Id: 3,
		Name: "asd",
		Description: "asd",
	}
	fmt.Println("main")
	json.NewEncoder(w).Encode(item)
}
