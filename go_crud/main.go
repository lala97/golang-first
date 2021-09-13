package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"go_crud/database"
	"go_crud/controllers"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	db_init()
	router := mux.NewRouter().StrictSlash(true)
	routerList(router)
	log.Fatal(http.ListenAndServe(":8090", router))

}


//routers list
func routerList(router *mux.Router) {
	router.HandleFunc("/", controllers.GetUserList).Methods("GET")
	router.HandleFunc("/user/{id}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/user/delete/{id}", controllers.DeletUserByID).Methods("DELETE")
	router.HandleFunc("/update/{id}", controllers.UpdateUser).Methods("PUT")

}

//initialize database and connect
func db_init() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go_crud",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
}
