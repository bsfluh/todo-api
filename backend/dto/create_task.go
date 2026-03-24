package dto

import (
	"time"
	"todo-Api/model"
)

type CreateTaskRequest struct {
	Title       string         `form:"title" binding:"required,min=1,max=200"`
	Description string         `form:"description" binding:"max=3000"`
	Priority    model.Priority `form:"priority" binding:"required,oneof=Paramount High Middle Low"`
	DeadLine    time.Time      `form:"deadline" binding:"required,gtfield=createdAt"`
}
