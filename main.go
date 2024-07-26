package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeysson/golang-api/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	router := mux.NewRouter()

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	handlers.Init(db)

	router.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	router.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", handlers.GetTodoByID).Methods("GET")
	router.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
