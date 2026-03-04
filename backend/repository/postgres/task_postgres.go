package postgres

import (
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
func (n *PoolTaskRepository) Create(task *model.Task) error {

}
func (n *PoolTaskRepository) GetByID(id int) (model.Task, error) {

}
func (n *PoolTaskRepository) GetAllByUserID(userID int) ([]model.Task, error) {

}
func (n *PoolTaskRepository) Update(task *model.Task) error {

}
func (n *PoolTaskRepository) Delete(id int, userId int) error {

}
