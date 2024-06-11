package models

type Client struct {
	ID         int    `json:"id"`
	Nome       string `json:"nome"`
	Idade      int    `json:"idade"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}
