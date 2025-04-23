package main

import (
	"flag"
	"github.com/Mrdeft2231/file-processing-api/tree/main/config"
	app "github.com/Mrdeft2231/file-processing-api/tree/main/internal/app"
	"log"
)

func main() {
	// Определяем флаг
	devMode := flag.Bool("dev", false, "Run server in development mode")
	flag.Parse()

	// Загружаем конфигурацию
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализируем логгер
	//appLogger := logger.New(*devMode)
	//defer appLogger.Sync()
	//// Настраиваем gRPC логгер
	//appLogger.ReplaceGrpcLogger()

	//	Запускаем
	app.Run(cfg, *devMode)
}
