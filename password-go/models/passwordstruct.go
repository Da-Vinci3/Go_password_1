package models

type Password struct {
	Service  string `json:"service"`
	Password string `json:"password"`
}
