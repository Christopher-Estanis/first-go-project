package handlers

import (
	"encoding/json"
	"example/models"
	"fmt"
	"log"
	"net/http"
)

func Create(res http.ResponseWriter, req *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(req.Body).Decode(&todo)

	if err != nil {
		log.Printf("Erro ao fazer decode do JSON: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Create(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserido com sucesso!"),
			"Id":      fmt.Sprintf("%d", id),
		}
	}

	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
