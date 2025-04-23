package grpcfile

import (
	"context"
	pg "github.com/Mrdeft2231/file-processing-api/tree/main/gen/file/proto"
	usecase "github.com/Mrdeft2231/file-processing-api/tree/main/internal/usecase/file"
)

var _ pg.FileProcessingServer = (*FileProcessingServer)(nil)

type FileProcessingServer struct {
	pg.UnimplementedFileProcessingServer
	file usecase.FileUseCase
}

func NewFileProcessingServer(file usecase.FileUseCase) *FileProcessingServer {
	return &FileProcessingServer{file: file}
}

func (s *FileProcessingServer) GetFiles(c context.Context, req *pg.GetFileRequest) (*pg.GetFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) UploadFile(ctx context.Context, req *pg.UploadFileRequest) (*pg.UploadFileResponse, error) {

	result, err := s.usecase.UploadFile(ctx, req)
	return nil, nil
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
