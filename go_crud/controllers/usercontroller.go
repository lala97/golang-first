package controllers

import (
	"encoding/json"
	"go_crud/database"
	"go_crud/model"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"fmt"
	"io/ioutil"
)
//get all users
func GetUserList(w http.ResponseWriter, r *http.Request) {
	var user []model.User
	database.Connector.Find(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

//Get user by spesific Id
func GetUserById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id :=vars["id"]
	var user model.User
	database.Connector.Find(&user,id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//DeletuserByID delete's user with specific ID
func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	fmt.Println(key)
	var user model.User
	id, _ := strconv.ParseInt(key, 10, 64)
	fmt.Println(id)

	res :=database.Connector.Where("id = ?", id).Delete(&user)
	fmt.Printf("%+v",res)

	// w.WriteHeader(http.StatusNoContent)
}

//Update user
func UpdateUser(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("update")
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person model.User
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}