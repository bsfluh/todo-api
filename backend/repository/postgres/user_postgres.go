package postgres

import (
	"todo-Api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DBUserRepository struct {
	pool *pgxpool.Pool
}

func NewDBUserRepository(pool *pgxpool.Pool) *DBUserRepository {
	return &DBUserRepository{
		pool: pool,
	}
}
func (m *DBUserRepository) Create(user *model.User) error {

}
func (m *DBUserRepository) GetByEmail(email string) (model.User, error) {

}
