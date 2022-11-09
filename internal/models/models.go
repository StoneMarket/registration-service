package models

type User struct {
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
