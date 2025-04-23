package grpcfile

import (
	"context"
	pg "github.com/Mrdeft2231/file-processing-api/tree/main/gen/file/proto"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/entity"
	usecase "github.com/Mrdeft2231/file-processing-api/tree/main/internal/usecase/file"
	"github.com/google/uuid"
	"github.com/h2non/filetype"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"time"
)

var _ pg.FileProcessingServer = (*FileProcessingServer)(nil)

type FileProcessingServer struct {
	pg.UnimplementedFileProcessingServer
	file usecase.FileUseCase
}

func NewFileProcessingServer(file usecase.FileUseCase) *FileProcessingServer {
	return &FileProcessingServer{file: file}
}

func (s *FileProcessingServer) GetFiles(c context.Context, req *pg.GetFilesRequest) (*pg.GetFilesResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) UploadFile(stream pg.FileProcessing_UploadFileServer) error {
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
			return status.Errorf(codes.Internal, "ошибка чтения чанка: %v", err)
		}

		if filename == "" && req.GetFilename() != "" {
			filename = req.GetFilename()
		}

		fullFile = append(fullFile, req.GetChunk()...)
	}

	if err := s.usecase.ProcessAndSave(fullFile, filename); err != nil {
		return status.Errorf(codes.Internal, "ошибка обработки файла: %v", err)
	}

	return stream.SendAndClose(&pb.UploadFileResponse{
		Message: "Файл успешно загружен и сохранён",
	})
}

func (s *FileProcessingServer) SearchFile(ctx context.Context, req *pg.SearchFileRequest) (*pg.SearchFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) ConvertFile(ctx context.Context, req *pg.ConvertFileRequest) (*pg.ConvertFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) DeleteFile(ctx context.Context, req *pg.DeleteFileRequest) (*pg.DeleteFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) AnalyzeFile(ctx context.Context, req *pg.AnalyzeFileRequest) (*pg.AnalyzeFileResponse, error) {
	return nil, nil
}
