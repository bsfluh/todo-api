package repository

import "todo-Api/model"

type UserRepository interface {
	Create(user *model.User) error
	GetByEmail(email string) (*model.User, error)
}
type TaskRepository interface {
	Create(task *model.Task) error
	GetByID(id int) (*model.Task, error)
	GetAllTasksByUserID(userID int) ([]model.Task, error)
	Update(task *model.Task) error
	Delete(id int, userID int) error
}
