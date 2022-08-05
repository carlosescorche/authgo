package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/carlosescorche/authgo/handlers"
	"github.com/carlosescorche/authgo/middlewares"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", handlers.HandlerUserLogin).Methods("POST")
	r.HandleFunc("/user", handlers.HandlerUserAdd).Methods("POST")

	current := r.PathPrefix("/user/current").Subrouter()
	current.Use(middlewares.MiddlewareAuth)
	current.HandleFunc("", handlers.HandlerUserGet).Methods("GET")
	current.HandleFunc("", handlers.HandlerUserUpdate).Methods("PUT")

	err := http.ListenAndServe(os.Getenv("port"), r)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Conectado al puerto ", os.Getenv("port"))
}
