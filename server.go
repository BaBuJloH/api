package api

import (
	"context"
	"net/http"
	"time"
)

type Server struct { //абстракция над стандартной структурой Server определяющая параметры запуска
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error { //метод для запуска сервера
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error { //метод для выхода из приложения
	return s.httpServer.Shutdown(ctx)
}
