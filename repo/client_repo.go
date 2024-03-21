package repo

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
