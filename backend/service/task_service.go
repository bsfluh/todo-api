package service

import (
	"todo-Api/model"
	"todo-Api/repository"
)

type TaskService interface {
	CreateTask(userID int, title, description string) (error, *model.Task)
}
type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: repo,
	}
}
func (s *taskService) CreateTask(userID int, title, description string) (error, *model.Task) {
	_, err := c.Cookie("user_id")
	task := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
	}
	err = s.taskRepo.CreateTask(task)
	return task, nil
}
