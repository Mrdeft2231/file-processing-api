package grpcfile

import (
	"context"
	pb "github.com/Mrdeft2231/file-processing-api/tree/main/gen/file/proto"
	usecase "github.com/Mrdeft2231/file-processing-api/tree/main/internal/usecase/file"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

var _ pb.FileProcessingServer = (*FileProcessingServer)(nil)

type FileProcessingServer struct {
	pb.UnimplementedFileProcessingServer
	file usecase.FileUseCase
}

func NewFileProcessingServer(file usecase.FileUseCase) *FileProcessingServer {
	return &FileProcessingServer{file: file}
}

func (s *FileProcessingServer) GetFiles(c context.Context, req *pb.GetFilesRequest) (*pb.GetFilesResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) UploadFile(stream pb.FileProcessing_UploadFileServer) error {
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

	if err := s.file.UploadFile(fullFile, filename); err != nil {
		return status.Errorf(codes.Internal, "ошибка обработки файла: %v", err)
	}

	return stream.SendAndClose(&pb.UploadFileResponse{
		Message: "Файл успешно загружен и сохранён",
	})
}

func (s *FileProcessingServer) SearchFile(ctx context.Context, req *pb.SearchFileRequest) (*pb.SearchFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) ConvertFile(ctx context.Context, req *pb.ConvertFileRequest) (*pb.ConvertFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.DeleteFileResponse, error) {
	return nil, nil
}

func (s *FileProcessingServer) AnalyzeFile(ctx context.Context, req *pb.AnalyzeFileRequest) (*pb.AnalyzeFileResponse, error) {
	return nil, nil
}
