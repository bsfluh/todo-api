package model

import "time"

type Priority string

const (
	Paramount Priority = "Paramount" //Срочные и важные
	High      Priority = "High"      //срочные и не важные
	Middle    Priority = "Middle"    //несрочные и важные
	Low       Priority = "Low"       //несрочные и не важные
)

type User struct {
	ID           int64     `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	CreatedAt    time.Time `json:"createdAt"`
}
type Task struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"userID"`
	Title       string    `json:"title"`
	Priority    Priority  `json:"priority"`
	Deadline    time.Time `json:"deadline"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
	Description string    `json:"description"`
}
