package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/auth-api/internal/entity"
)

// Убедимся, что repository реализует интерфейс Repository
var _ Repository = (*repository)(nil)

type Repository interface {
	// CreateUser - создание нового пользователя
	CreateUser(ctx context.Context, user *entity.User) (string, error)
	// GetUserByUsername - получение пользователя по email
	GetUserByUsername(ctx context.Context, email string) (*entity.User, error)
	// GetUserById - получение пользователя по id
	GetUserById(ctx context.Context, id string) (*entity.User, error)
}

// repository - репозиторий для работы с PostgreSQL
type repository struct {
	db *pgxpool.Pool
}

// NewRepository - конструктор создания репозитория для работы с PostgreSQL
func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *entity.User) (string, error) {
	query := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(ctx, query, user.ID, user.Username, user.Password)
	if err != nil {
		return "", err
	}

	return user.ID, nil
}

func (r *repository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT id, username, password, is_active FROM users WHERE username = $1`
	row := r.db.QueryRow(ctx, query, username)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *repository) GetUserById(ctx context.Context, id string) (*entity.User, error) {
	query := `SELECT id, username, password, is_active FROM users WHERE id = $1`
	row := r.db.QueryRow(ctx, query, id)

	var user entity.User

	if err := row.Scan(&user.ID, &user.Username, &user.Password, &user.IsActive); err != nil {
		return nil, err
	}
	return &user, nil
}
