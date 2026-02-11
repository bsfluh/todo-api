package model

import "time"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
}
type Task struct {
	ID         int       `json:"id"`
	UserID     int       `json:"userID"`
	Title      string    `json:"title"`
	CreadTime  time.Time `json:"creadTime"`
	UpdateTime time.Time `json:"updateTime"`
	Deadline   time.Time `json:"dedline"`
}
type TypesTask struct {
	ParamountImportance Task `json:"paramountImportance"`
	HighImportance      Task `json:"highImportance"`
	MiddlImportance     Task `json:"middlImportance"`
	LowImportance       Task `json:"lowImportance"`
}
