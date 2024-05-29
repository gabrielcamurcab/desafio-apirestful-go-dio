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
	r := mux.NewRouter()
	r.HandleFunc("/client", handlers.CreateClient).Methods("POST")
	r.HandleFunc("/client", handlers.GetClients).Methods("GET")
	r.HandleFunc("/client/{id}", handlers.GetClientById).Methods("GET")
	r.HandleFunc("/client/{id}", handlers.DeleteClientById).Methods("DELETE")
	r.HandleFunc("/client/{id}", handlers.UpdateClient).Methods("PUT")
	return r
}
