package usecase

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/entity"
	repository "github.com/Mrdeft2231/file-processing-api/tree/main/internal/repo/file"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"time"
	"unicode"
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
	fmt.Println("в сервисе")

	king, err := filetype.Match(file)
	fmt.Printf("king %+v\n", king)
	if err != nil || king == filetype.Unknown {
		reader := bytes.NewReader(file)
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Text()

			for _, r := range line {
				if r > unicode.MaxASCII {
					return errors.New("файл не опознан")
				}
			}
			king.MIME.Value = "text/.txt"
			king.Extension = ".txt"
		}
	}

	userID := uuid.New().String()
	fmt.Printf("service %+v\n", king)
	f := &entity.File{
		FileID:    userID,
		Name:      filename,
		MimeType:  king.MIME.Value,
		Extension: king.Extension,
		Size:      len(file),
		Content:   file,
		CreatedAt: time.Now(),
	}
	fmt.Printf("f %+v\n", f)
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
