package main

import (
	"log"
	"net/http"

	"github.com/jeysson/golang-api/config"
	"github.com/jeysson/golang-api/routes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	loadPort, err2 := config.LoadConfig("app.json")

	if err2 != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err2)
	}

	// criação da instância de App
	// Creating App instance
	app := &config.App{DB: db}

	// Inicialização das rotas
	// Routes initializing
	router := routes.InitRoutes(app)

	log.Printf("Iniciando o servidor na porta %s...\n", loadPort.Port)
	if err := http.ListenAndServe(":"+loadPort.Port, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v\n", err)
	}
}
