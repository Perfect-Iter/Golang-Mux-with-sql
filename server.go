package main

import (
	"goMux/config"
	controller "goMux/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":8080"
)

func initializeRouter() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controller.GetUsers).Methods("GET")
	router.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/users", controller.CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	log.Println("Server starting on port: ", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	config.InitialMigration()
	initializeRouter()

}
