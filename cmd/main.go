package main

import (
	"api"
	"api/pkg/handler"
	"api/pkg/repository"
	"api/pkg/service"
	"log"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("ошибка в инициализации конфигурации: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("ошибка инициализации подключения к бд: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server) //инициализация экземпляра сервера
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	} // запуск сервера с помощью метода Run

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
