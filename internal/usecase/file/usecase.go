package usecase

import (
	repository "github.com/Mrdeft2231/file-processing-api/tree/main/internal/repo/file"
	"github.com/golang-migrate/migrate/v4/source/file"
)

var _ FileUseCase = (*fileUseCase)(nil)

type FileUseCase interface {
	GetFiles() (*file.File, error)
	UploadFile() error
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

func (uc *fileUseCase) UploadFile() error {
	return nil
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
