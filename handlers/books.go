package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeysson/golang-api/config"
	"github.com/jeysson/golang-api/models"
)

type HandlerBook struct {
	App *config.App
}

func (hb *HandlerBook) GetBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	hb.App.DB.Preload("Authors").Find(&books)
	json.NewEncoder(w).Encode(books)
}

func (hb *HandlerBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)
	hb.App.DB.Create(&book)
	json.NewEncoder(w).Encode(book)
}

func (hb *HandlerBook) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	hb.App.DB.Preload("Authors").First(&book, params["id"])
	json.NewEncoder(w).Encode(book)
}

func (hb *HandlerBook) UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	hb.App.DB.First(&book, params["id"])
	json.NewDecoder(r.Body).Decode(&book)
	hb.App.DB.Save(&book)
	json.NewEncoder(w).Encode(book)
}

func (hb *HandlerBook) DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	hb.App.DB.Delete(&book, params["id"])
	json.NewEncoder(w).Encode("Book deleted!")
}
