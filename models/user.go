package models

type User struct {
	Id         int    `json:"id"`
	Usuario    string `json:"usuario"`
	Senha      string `json:"senha"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
