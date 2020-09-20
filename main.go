package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Get("/", func(response http.ResponseWriter, request *http.Request) {
		response.Write([]byte("welcome"))
	})
	http.ListenAndServe(":3000", router)
}