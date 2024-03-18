package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"minhaapi/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Client representa um cliente

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	// Iniciar o servidor HTTP
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	log.Fatal(http.ListenAndServe(serverAddr, router()))
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cliente", handlers.CreateClient).Methods("POST")
	r.HandleFunc("/cliente", handlers.GetClients).Methods("GET")
	r.HandleFunc("/cliente/{id}", handlers.GetClientById).Methods("GET")
	return r
}
