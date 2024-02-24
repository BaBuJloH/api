package main

import (
	"api"
	"api/pkg/handler"
	"api/pkg/repository"
	"api/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(api.Server) //инициализация экземпляра сервера
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Ошибка запуска сервера: %s", err.Error())
	} // запуск сервера с помощью метода Run

}
