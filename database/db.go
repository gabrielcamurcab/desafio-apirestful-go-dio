package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

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
