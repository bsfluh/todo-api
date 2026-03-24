package service

import (
	"time"
	"todo-Api/model"
	"todo-Api/repository"
)

type TaskService interface {
	CreateTask(userID int64, title, description string, priority model.Priority, deadline, createdAt, updateAt time.Time) (error, *model.Task)
}
type taskService struct {
	taskRepo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		taskRepo: repo,
	}
}
func (s *taskService) CreateTask(userID int64, title, description string, priority model.Priority, deadline, createdAt, updateAt time.Time) (error, *model.Task) {
	task := &model.Task{
		UserID:      userID,
		Title:       title,
		Description: description,
		Priority:    priority,
		Deadline:    deadline,
		CreatedAt:   createdAt,
		UpdateAt:    updateAt,
	}
	err := s.taskRepo.CreateTask(task)
	if err != nil {
		return err, nil
	}
	return nil, task
}
