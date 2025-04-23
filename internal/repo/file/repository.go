package repository

import (
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/entity"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
)

var _ Repository = (*repository)(nil)

type Repository interface {
	GetFiles() (*file.File, error)
	UploadFile(file *entity.File) error
	DeleteFile() error
	SearchFile() (*file.File, error)
	ConvertFile() (*file.File, error)
	AnalyzeFile() (*file.File, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repository{db: db}
}

func (r *repository) GetFiles() (*file.File, error) {
	return nil, nil
}

func (r *repository) UploadFile(file *entity.File) error {
	query := `INSERT INTO files (name, size, mime_type, extension, content, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := r.db.Exec(context.Background(), query, file.Name, file.Size, file.MimeType, file.Extension, file.Content, file.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) DeleteFile() error {
	return nil
}

func (r *repository) SearchFile() (*file.File, error) {
	return nil, nil
}

func (r *repository) ConvertFile() (*file.File, error) {
	return nil, nil
}

func (r *repository) AnalyzeFile() (*file.File, error) {
	return nil, nil
}
