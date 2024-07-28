package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jeysson/golang-api/config"
	"github.com/jeysson/golang-api/models"
)

// o App só pode ser acesso por meio do encapsulamento
// por isso foi declarado esse struct
type HandlerAuthor struct {
	App *config.App
}

func (ha *HandlerAuthor) GetAuthors(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author
	ha.App.DB.Preload("Books").Find(&authors)
	json.NewEncoder(w).Encode(authors)
}

func (ha *HandlerAuthor) CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var author models.Author
	json.NewDecoder(r.Body).Decode(&author)
	ha.App.DB.Create(&author)
	json.NewEncoder(w).Encode(author)
}

func (ha *HandlerAuthor) GetAuthorByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	ha.App.DB.Preload("Books").Find(&author, params["id"])
}

func (ha *HandlerAuthor) UpdateAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	ha.App.DB.First(&author, params["id"])
	json.NewDecoder(r.Body).Decode(&author)
	ha.App.DB.Save(&author)
	json.NewEncoder(w).Encode(author)
}

func (hb *HandlerAuthor) DeleteAuthor(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var author models.Author
	hb.App.DB.Delete(&author, params["id"])
	json.NewEncoder(w).Encode("Author deleted!")
}
