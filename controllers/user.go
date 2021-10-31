package controller

import (
	"encoding/json"
	"goMux/config"
	"goMux/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var users []models.User

	config.DB.Find(&users)

	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	var user models.User

	config.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var user models.User
	params := mux.Vars(r)

	if config.DB.First(&user, params["email"]) != nil {
		log.Println("User already exists")
	} else {
		json.NewDecoder(r.Body).Decode(&user)

		config.DB.Create(&user)
		json.NewEncoder(w).Encode(user)
	}

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	var user models.User

	config.DB.First(&user, params["id"])

	json.NewDecoder(r.Body).Decode(&user)
	config.DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	params := mux.Vars(r)
	var user models.User

	config.DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}
