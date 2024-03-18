package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Client representa um cliente
type Client struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

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

// ConnectDB cria e retorna uma conexão com o banco de dados
func ConnectDB() (*sql.DB, error) {
	// Obter informações de conexão do arquivo .env
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construir string de conexão
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Abrir conexão com o banco de dados
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cliente", createClient).Methods("POST") // Rota para adicionar um novo cliente
	return r
}

func createClient(w http.ResponseWriter, r *http.Request) {
	// Decodificar o corpo da solicitação em um struct Client
	var newClient Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados
	db, err := ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Preparar a declaração SQL para inserir o novo cliente na tabela "cliente"
	stmt, err := db.Prepare("INSERT INTO clientes(nome, idade) VALUES(?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	// Executar a declaração SQL para adicionar o novo cliente
	result, err := stmt.Exec(newClient.Nome, newClient.Idade)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obter o ID do novo cliente adicionado
	newID, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Definir o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")
	// Retornar o ID do novo cliente como resposta
	json.NewEncoder(w).Encode(map[string]string{"message": "Registro criado com sucesso!", "id": strconv.FormatInt(newID, 10)})
}
