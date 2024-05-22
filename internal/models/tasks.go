package models

import "time"

type Tasks struct {
	Id         int        `json:"id" db:"id"`
	Opened     *time.Time `json:"opened" db:"opened"`
	Closed     *time.Time `json:"closed" db:"closed"`
	AuthorId   int        `json:"author_id" db:"author_id"`
	AssignedId int        `json:"assigned_id" db:"assigned_id"`
	Title      string     `json:"title" db:"title"`
	Content    string     `json:"content" db:"content"`
}

type CreateTask struct {
	AuthorId   int    `json:"author_id" db:"author_id"`
	AssignedId int    `json:"assigned_id" db:"assigned_id"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
}

type UpdateTask struct {
	Id         int    `json:"id" db:"id"`
	AuthorId   int    `json:"author_id" db:"author_id"`
	AssignedId int    `json:"assigned_id" db:"assigned_id"`
	Title      string `json:"title" db:"title"`
	Content    string `json:"content" db:"content"`
}
