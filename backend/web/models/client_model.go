package models

type Client struct {
	ClientID  int    `json:"clientID"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	password  string `json:"password"`
}
