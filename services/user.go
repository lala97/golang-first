package services

import (
	"encoding/json"
	 "go_crud/models"
	"go_crud/responses"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//get all users
func GetUserList(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	users, err := user.FindAll()
	if err != nil {
		responses.ErrorJson(w, http.StatusInternalServerError, err)
	}
	responses.SuccessJson(w, http.StatusOK, users)
}

//Get user by spesific Id
func GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	user := model.User{}
	data, err := user.FindById(int32(id))
	if err != nil {
		responses.ErrorJson(w, http.StatusInternalServerError, err)
	}
	responses.SuccessJson(w, http.StatusFound, data)
}

//DeletuserByID delete's user with specific ID
func DeletUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	user := model.User{}
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	err := user.Delete(int32(id))
	if err != nil {
		responses.ErrorJson(w, http.StatusInternalServerError, err)
	}
	responses.SuccessJson(w, http.StatusNoContent, "")

}

//Create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	user := model.User{}
	json.Unmarshal(requestBody, &user)
	newUser, err := user.Create(&user)
	if err != nil {
		responses.ErrorJson(w, http.StatusInternalServerError, err)
	}
	responses.SuccessJson(w, http.StatusCreated, newUser)
}

//Update user information
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 32)
	user := model.User{}
	json.Unmarshal(requestBody, &user)
	data, err := user.Update(int32(id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
	responses.SuccessJson(w, http.StatusOK, data)
}
