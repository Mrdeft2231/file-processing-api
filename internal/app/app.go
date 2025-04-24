// Package app configures and runs application.
package app

import (
	"context"
	"fmt"
	"github.com/Mrdeft2231/file-processing-api/tree/main/config"
	file "github.com/Mrdeft2231/file-processing-api/tree/main/gen/file/proto"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/adapter/postgres"
	grpcfile "github.com/Mrdeft2231/file-processing-api/tree/main/internal/controller/grpc/file"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/repo/file"
	"github.com/Mrdeft2231/file-processing-api/tree/main/internal/usecase/file"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"net"
)

// Run - запускает приложение
func Run(cfg *config.Config, devMode bool) {
	// Инициализация дефолтного логгера
	logger := log.Default()
	// Подключение к базе данных
	dbpool, err := postgres.New(context.Background(), *cfg)
	if err != nil {
		logger.Fatalf("Unable to create connection pool: %v", err)
	}
	defer dbpool.Close()
	logger.Printf("Database connection established")

	// Создаем репозитории
	fileRepo := repository.NewRepository(dbpool)

	// Создаем слой usecase
	fileUseCase := usecase.NewFileUseCase(fileRepo)

	// Создаем gRPC-сервер
	grpcServer := grpc.NewServer(
		grpc.StreamInterceptor(grpcLogStreamInterceptor),
		grpc.UnaryInterceptor(grpcLogUnaryInterceptor),
	)

	// Создаем и регистрируем gRPC-сервис
	fileController := grpcfile.NewFileProcessingServer(fileUseCase)
	file.RegisterFileProcessingServer(grpcServer, fileController)

	// Слушаем порт gRPC
	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", cfg.GRPC.Port))
	if err != nil {
		logger.Fatalf("Failed to listen on port %d: %v", cfg.GRPC.Port, err)
	}

	logger.Printf("Starting gRPC server on port %d\n", cfg.GRPC.Port)
	// Запускаем gRPC-сервер
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatalf("Failed to serve gRPC server: %v", err)
	}

}

// Интерсепторы для логирования
func grpcLogStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger := log.Default()

	if p, ok := peer.FromContext(ss.Context()); ok && p.Addr != nil {
		logger.Printf("gRPC Stream called: %s from %s", info.FullMethod, p.Addr.String())
	} else {
		logger.Printf("gRPC Stream called: %s from UNKNOWN", info.FullMethod)
	}

	return handler(srv, ss)
}

func grpcLogUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	logger := log.Default()

	p, ok := peer.FromContext(ctx)
	addr := "UNKNOWN"
	if ok && p.Addr != nil {
		addr = p.Addr.String()
	}

	resp, err := handler(ctx, req)

	if err != nil {
		logger.Printf("gRPC Unary response error: %s, method: %s, error: %v", addr, info.FullMethod, err)
	} else {
		logger.Printf("gRPC Unary response: %s, method: %s, response: %v", addr, info.FullMethod, resp)
	}

	return resp, err
}
