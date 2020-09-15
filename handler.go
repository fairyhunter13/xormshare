package main

import (
	"encoding/json"
	"fmt"
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

type errResp struct {
	Error string `json:"error"`
}

func printErr(w http.ResponseWriter, err error, status int) {
	w.WriteHeader(status)
	var resp errResp
	resp.Error = err.Error()
	json.NewEncoder(w).Encode(&resp)
}

func printSuccess(w http.ResponseWriter, payload interface{}, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		printErr(w, err, http.StatusBadRequest)
		return
	}
	session := engine.Prepare()
	defer session.Close()
	var isExist bool
	isExist, err = session.Get(&user)
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}
	if !isExist {
		printErr(w, fmt.Errorf("Error user not found, user struct: %+v", user), http.StatusNotFound)
		return
	}

	printSuccess(w, user, http.StatusOK)
}

func addUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		printErr(w, err, http.StatusBadRequest)
		return
	}
	session := engine.Prepare()
	defer session.Close()
	_, err = engine.InsertOne(&user)
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}

	printSuccess(w, user, http.StatusCreated)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var (
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		printErr(w, err, http.StatusBadRequest)
		return
	}

	session := engine.Prepare()
	defer session.Close()
	id := user.ID
	user.ID = 0
	_, err = session.ID(id).Update(&user)
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}
	var (
		updatedUser User
		isExist     bool
	)
	updatedUser.ID = id
	isExist, err = session.Get(&updatedUser)
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}
	if !isExist {
		printErr(w, fmt.Errorf("Error user not found, user struct: %+v", updatedUser), http.StatusNotFound)
		return
	}

	printSuccess(w, updatedUser, http.StatusOK)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var user User
	var (
		err error
	)
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		printErr(w, err, http.StatusBadRequest)
		return
	}

	session := engine.Prepare()
	defer session.Close()

	var isExist bool
	isExist, err = session.Get(&user)
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}
	if !isExist {
		printErr(w, fmt.Errorf("Error user not found, user struct: %+v", user), http.StatusNotFound)
		return
	}

	_, err = session.ID(user.ID).Delete(&User{})
	if err != nil {
		printErr(w, err, http.StatusInternalServerError)
		return
	}

	printSuccess(w, user, http.StatusOK)
}
