package main

import (
	"fmt"
	"log"
	"net/http"

	"Latihan1/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", controllers.GetAllUsers).Methods("GET")
	router.HandleFunc("/users", controllers.Authenticate(controllers.InsertUser, 1)).Methods("POST")
	router.HandleFunc("/users/{user_id}", controllers.Authenticate(controllers.DeleteUser, 1)).Methods("DELETE")
	router.HandleFunc("/login", controllers.Login).Methods("GET")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
