package models

type User struct {
	Image    []byte `json:"-"`
	ID       int    `json:"-"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Response struct {
	Type        string `json:"type"`
	AccessLogin string `json:"access_login"`
}
