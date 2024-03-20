package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"minhaapi/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Servidor iniciado na porta %s\n", port)
	log.Fatal(http.ListenAndServe(serverAddr, router()))
}

func router() http.Handler {
	r := http.NewServeMux()
	r.HandleFunc("/cliente/store", handlers.CreateClient)
	r.HandleFunc("/cliente/get", handlers.GetClients)
	r.HandleFunc("/cliente/get/{id}", handlers.GetClientById)
	r.HandleFunc("/cliente/drop/{id}", handlers.DeleteClientById)
	r.HandleFunc("/cliente/update/{id}", handlers.UpdateClient)
	return r
}
