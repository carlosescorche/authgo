package main

import (
	"log"
	"net/http"
	"os"

	"github.com/carlosescorche/authgo/handlers"
	"github.com/carlosescorche/authgo/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/login", handlers.HandlerUserLogin).Methods("POST")
	r.HandleFunc("/user", handlers.HandlerUserAdd).Methods("POST")
	//r.HandleFunc("/verify", handlers.HandlerUserVerify).Methods("PUT")

	current := r.PathPrefix("/user/current").Subrouter()
	current.Use(middlewares.MiddlewareAuth)
	current.HandleFunc("", handlers.HandlerUserGet).Methods("GET")
	//r.HandleFunc("/", handlers.HandlerUserUpdate).Methods("PUT")
	//r.HandleFunc("/password", handlers.HandlerUserUpdatePassword).Methods("PUT")
	//r.HandleFunc("/logout", handlers.HandlerUserUpdate).Methods("PUT")

	err := http.ListenAndServe(os.Getenv("port"), r)

	if err != nil {
		log.Fatal(err.Error())
	}
}
