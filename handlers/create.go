package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Danielopes7/api-go/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error decoding the request body: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"message": fmt.Sprint("Error creating the todo: %v", err),
			"Error":   true,
		}
	} else {
		resp = map[string]any{
			"message": fmt.Sprint("Todo created successfully with id: %v", id),
			"Error":   false,
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
