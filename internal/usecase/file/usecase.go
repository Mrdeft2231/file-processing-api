package usecase

import (
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/entity"
	repository "github.com/Mrdeft2231/file-processing-api/tree/main/internal/repo/file"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
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
	var (
		fullFile []byte
		filename string
	)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return status.Errorf(codes.Internal, "failed to receive request: %v", err)
		}

		fullFile = append(fullFile, req.GetChunk()...)
	}

	king, err := filetype.Match(fullFile)
	if err != nil || king == filetype.Unknown {
		return status.Error(codes.InvalidArgument, "file type not supported")
	}

	userId := uuid.New().String()

	file := entity.File{
		FileID:    userId,
		Name:      filename,
		Size:      int64(len(fullFile)),
		MimeType:  king.MIME.Value,
		Extension: king.Extension,
		CreatedAt: time.Now(),
	}
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
