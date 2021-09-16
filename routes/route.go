package routes

import (
	"go_crud/services"
	"net/http"

	"github.com/gorilla/mux"
)

//routers list
func Initialize() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//user routes
	user := router.PathPrefix("/users").Subrouter()
	user.HandleFunc("/list", services.GetUserList).Methods("GET")
	user.HandleFunc("/{id}", services.GetUserById).Methods("GET")
	user.HandleFunc("/delete/{id}", services.DeletUserByID).Methods("DELETE")
	user.HandleFunc("/update/{id}", services.UpdateUser).Methods("PUT")
	user.HandleFunc("/create", services.CreateUser).Methods("POST")
	http.ListenAndServe(":8090", router)

	return router
}
