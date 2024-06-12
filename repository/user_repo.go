package repository

import (
	"database/sql"
	"fmt"
	"minhaapi/models"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *sql.DB, newUser models.User) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO usuarios(usuario, senha) VALUES(?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	senhaHash, err := bcrypt.GenerateFromPassword([]byte(newUser.Senha), 12)
	if err != nil {
		fmt.Println("Erro ao gerar hash de senha", err)
		return 0, err
	}

	result, err := stmt.Exec(newUser.Usuario, senhaHash)
	if err != nil {
		fmt.Println("Erro ao criar o usuário", err)
		return 0, err
	}

	return result.LastInsertId()
}
