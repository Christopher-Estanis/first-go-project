package main

import (
	"example/configs"
	"example/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()

	router.Post("/", handlers.Create)
	router.Get("/", handlers.List)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/{id}", handlers.Show)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)

}
