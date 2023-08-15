package handlers

import (
	"encoding/json"
	"example/models"
	"log"
	"net/http"
)

func List(res http.ResponseWriter, req *http.Request) {
	todos, err := models.List()

	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	res.Header().Add("Content-type", "application/json")
	json.NewEncoder(res).Encode(todos)
}
