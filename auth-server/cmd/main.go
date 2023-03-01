package main

import (
	"auth-server/api"
	"auth-server/api/handler"
	"auth-server/api/repository"
	"auth-server/api/service"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func initConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Configuration error:: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	_repository := repository.NewRepository(db)
	_service := service.NewService(_repository)
	_handler := handler.NewHandler(_service)

	srv := new(api.Server)
	if err := srv.Run(os.Getenv("GATEWAY_PORT"), _handler.InitHandler()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}
