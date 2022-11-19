package models

type User struct {
	Image    []byte `json:"-"`
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
