package main

import (
	"FinalProject/internal/config"
	"FinalProject/internal/handlers"
	"FinalProject/internal/repository"
	"FinalProject/internal/service"
	"FinalProject/pkg/logging"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"net/http"
)

func main() {
	configs, err := config.InitConfigs()
	if err != nil {
		log.Fatal("Failed to read configs: ", err)
	}

	logger, err := logging.InitializeLogger()
	if err != nil {
		log.Fatal("Failed to initialize logger: ", err)
	}

	db, err := connectToDB(configs.Database)
	if err != nil {
		log.Fatal("Failed to connect to database, err: ", err)
	}

	newRepository := repository.NewRepository(db, logger)
	newService := service.NewService(*newRepository)
	mux := gin.Default()
	newHandler := handlers.NewHandler(mux, newService, logger)
	newHandler.InitRoutes()
	serverAddr := fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port)
	http.ListenAndServe(serverAddr, newHandler)
}

func connectToDB(dbConf config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConf.Host,
		dbConf.Port,
		dbConf.User,
		dbConf.Password,
		dbConf.DBName,
		dbConf.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %v", err)
	}

	return db, nil
}
