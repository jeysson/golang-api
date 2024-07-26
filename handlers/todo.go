package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeysson/golang-api/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
	DB.AutoMigrate(&models.Todo{})
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	DB.First(&todo, params["id"])
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	DB.First(&todo, params["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	DB.Delete(&todo, params["id"])
	json.NewEncoder(w).Encode("Todo Deleted")
}
