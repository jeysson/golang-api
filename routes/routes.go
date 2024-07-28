package routes

import (
	"github.com/gorilla/mux"
	"github.com/jeysson/golang-api/config"
	"github.com/jeysson/golang-api/handlers"
	"github.com/jeysson/golang-api/middleware"
)

func InitRoutes(app *config.App) *mux.Router {
	router := mux.NewRouter()

	authHandler := &handlers.HandleAuth{App: app}
	authorHandler := &handlers.HandlerAuthor{App: app}
	booksHandler := &handlers.HandlerBook{App: app}

	// handler de autenticação
	router.HandleFunc("/register", authHandler.Register).Methods("POST")
	router.HandleFunc("/login", authHandler.Login(app.SecretKey)).Methods("POST")

	router.Use(middleware.LogginMiddleware)
	router.Use(middleware.RateLimiterMiddleware)

	// Rotas públicas sem autenticação
	router.HandleFunc("/authors", authorHandler.GetAuthors).Methods("GET")
	router.HandleFunc("/books", booksHandler.GetBooks).Methods("GET")

	// Rotas protegidas (precisam de autenticação JWT)
	protected := router.PathPrefix("/").Subrouter() // Criar um subrouter para rotas protegidas
	protected.Use(middleware.JWTMiddleware("your_secrete_key"))

	// Rotas para Authors
	protected.HandleFunc("/authors", authorHandler.CreateAuthor).Methods("POST")
	protected.HandleFunc("/authors/{id}", authorHandler.GetAuthorByID).Methods("GET")
	protected.HandleFunc("/authors/{id}", authorHandler.UpdateAuthor).Methods("PUT")
	protected.HandleFunc("/authors/{id}", authorHandler.DeleteAuthor).Methods("DELETE")

	// Rotas para Books
	protected.HandleFunc("/books", booksHandler.CreateBook).Methods("POST")
	protected.HandleFunc("/books/{id}", booksHandler.GetBookByID).Methods("GET")
	protected.HandleFunc("/books/{id}", booksHandler.UpdateBook).Methods("PUT")
	protected.HandleFunc("/books/{id}", booksHandler.DeleteBook).Methods("DELETE")

	return router
}
