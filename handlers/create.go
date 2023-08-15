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

	if err != nil {
		log.Printf("Ocorreu um erro ao tentar inserir: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Todo inserido com sucesso!"),
		"Id":      fmt.Sprintf("%d", id),
	}

	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(resp)
}
