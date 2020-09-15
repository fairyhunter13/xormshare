package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

var (
	mux chi.Router
)

func initHandler() {
	mux = chi.NewMux()
	mux.Route("/user", func(r chi.Router) {
		r.Get("/", getUser)
		r.Post("/", addUser)
		r.Patch("/", updateUser)
		r.Delete("/", deleteUser)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {

}

func addUser(w http.ResponseWriter, r *http.Request) {

}

func updateUser(w http.ResponseWriter, r *http.Request) {

}

func deleteUser(w http.ResponseWriter, r *http.Request) {

}
