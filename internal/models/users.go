package models

type User struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
