package postgres

import (
	"context"
	"todo-Api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolTaskRepository struct {
	pool *pgxpool.Pool
}

func NewPoolTaskRepository(pool *pgxpool.Pool) *PoolTaskRepository {
	return &PoolTaskRepository{
		pool: pool,
	}
}

func (n *PoolTaskRepository) CreateTask(task *model.Task) error {
	query := `
INSERT INTO tasks (user_id, title, priority, deadline, created_at, update_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id
`
	err := n.pool.QueryRow(context.Background(), query, task.UserID, task.Title, task.Priority, task.Deadline, task.CreatedAt, task.UpdateAt).Scan(&task.ID)
	if err != nil {
		return err
	}
	return nil

}

/*
func (n *PoolTaskRepository) GetByID(id int) (model.Task, error) {

}
func (n *PoolTaskRepository) GetAllByUserID(userID int) ([]model.Task, error) {

}
func (n *PoolTaskRepository) Update(task *model.Task) error {

}
func (n *PoolTaskRepository) Delete(id int, userId int) error {

}
*/
