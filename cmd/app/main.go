package main

import (
	"flag"
	"gitlab.crja72.ru/golang/2025/spring/course/projects/go21/auth-api/config"
	app "gitlab.crja72.ru/golang/2025/spring/course/projects/go21/auth-api/internal/app"
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
