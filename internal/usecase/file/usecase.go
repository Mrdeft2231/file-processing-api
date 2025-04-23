package usecase

import (
	"errors"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/entity"
	repository "github.com/Mrdeft2231/file-processing-api/tree/main/internal/repo/file"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"time"
)

var _ FileUseCase = (*fileUseCase)(nil)

type FileUseCase interface {
	GetFiles() (*file.File, error)
	UploadFile(file []byte, filename string) error
	DeleteFile() error
	SearchFile() (*file.File, error)
	ConvertFile() (*file.File, error)
	AnalyzeFile() (*file.File, error)
}

type fileUseCase struct {
	fileRepo repository.Repository
}

func NewFileUseCase(fileRepo repository.Repository) FileUseCase {
	return &fileUseCase{fileRepo: fileRepo}
}

func (uc *fileUseCase) GetFiles() (*file.File, error) {
	return nil, nil
}

func (uc *fileUseCase) UploadFile(file []byte, filename string) error {
	king, err := filetype.Match(file)
	if err != nil || king == filetype.Unknown {
		return errors.New("unknown file type")
	}

	userID := uuid.New().String()

	f := &entity.File{
		FileID:    userID,
		Name:      filename,
		MimeType:  king.MIME.Value,
		Extension: king.Extension,
		Content:   file,
		CreatedAt: time.Now(),
	}
	return uc.fileRepo.UploadFile(f)
}

func (uc *fileUseCase) DeleteFile() error {
	return nil
}

func (uc *fileUseCase) SearchFile() (*file.File, error) {
	return nil, nil
}

func (uc *fileUseCase) ConvertFile() (*file.File, error) {
	return nil, nil
}

func (uc *fileUseCase) AnalyzeFile() (*file.File, error) {
	return nil, nil
}
