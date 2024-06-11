package repository

import (
	"database/sql"
	"minhaapi/models"
)

func CreateClient(db *sql.DB, newClient models.Client) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO clientes(nome, idade) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(newClient.Nome, newClient.Idade)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func GetClients(db *sql.DB) ([]models.Client, error) {
	rows, err := db.Query("SELECT * FROM clientes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	clients := []models.Client{}
	for rows.Next() {
		var cliente models.Client
		if err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Idade, &cliente.Created_at, &cliente.Updated_at); err != nil {
			return nil, err
		}
		clients = append(clients, cliente)
	}

	return clients, nil
}

func GetClientById(db *sql.DB, id string) ([]models.Client, error) {
	stmt, err := db.Prepare("SELECT * FROM clientes WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var clientes []models.Client

	for rows.Next() {
		var cliente models.Client
		err := rows.Scan(&cliente.ID, &cliente.Nome, &cliente.Idade, &cliente.Created_at, &cliente.Updated_at)
		if err != nil {
			return nil, err
		}
		clientes = append(clientes, cliente)
	}

	return clientes, nil
}
