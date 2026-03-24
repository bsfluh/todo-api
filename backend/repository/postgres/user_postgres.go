package postgres

import (
	"context"
	"todo-Api/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PoolUserRepository struct {
	pool *pgxpool.Pool
}

func NewPoolUserRepository(pool *pgxpool.Pool) *PoolUserRepository {
	return &PoolUserRepository{
		pool: pool,
	}
}
func (m *PoolUserRepository) Create(user *model.User) error {
	query := `
INSERT INTO users( name, email, password_hash, created_at)
VALUES ($1, $2, $3, $4)
RETURNING id
`
	err := m.pool.QueryRow(context.Background(), query, user.Name, user.Email, user.PasswordHash, user.CreatedAt).Scan(&user.ID)
	return err
}

func (m *PoolUserRepository) GetByEmail(email string) (*model.User, error) {
	query := `
SELECT id, name, email, password_hash, created_at
FROM users
WHERE email = $1
`
	user := &model.User{}
	err := m.pool.QueryRow(context.Background(), query, email).Scan(&user.ID, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
