package handlers

import (
	"encoding/json"
	"minhaapi/database"
	"minhaapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateClient(w http.ResponseWriter, r *http.Request) {
	// Decodificar o corpo da solicitação em um struct Client
	var newClient models.Client
	err := json.NewDecoder(r.Body).Decode(&newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Conectar ao banco de dados
	db, err := database.ConnectDB()
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

func GetClients(w http.ResponseWriter, r *http.Request) {
	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	clients := []models.Client{}
	for rows.Next() {
		var client models.Client
		if err := rows.Scan(&client.ID, &client.Nome, &client.Idade); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clients = append(clients, client)
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(clients); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetClientById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT * FROM clientes WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(clientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var clientes []models.Client

	for rows.Next() {
		var cliente models.Client
		err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Idade)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(clientes[0]); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteClientById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	clientId := params["id"]

	db, err := database.ConnectDB()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM clientes WHERE id = ?")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Query(clientId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer result.Close()

	w.Header().Set("Content-type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "Registro excluído com sucesso!"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}