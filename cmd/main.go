package main

import (
	"api"
	"api/pkg/handler"
	"api/pkg/repository"
	"api/pkg/service"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка в инициализации конфигурации: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка загрузки переменных окружения: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("ошибка инициализации подключения к бд: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server) //инициализация экземпляра сервера

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Ошибка запуска сервера: %s", err.Error())
		} // запуск сервера с помощью метода Run
	}()
	logrus.Print("Приложение api запустилось")

	quit := make(chan os.Signal, 1)                      //канал для блокировки main функции
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT) //запись в канал происходит при получении сигналов системы
	<-quit

	logrus.Print("api Shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("ошибка закрытия соединения с сервером: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("ошибка закрытия соединения с базой данных: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
