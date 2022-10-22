package http

import (
	"net/http"
	"user_balance/service/internal/helpers"
	"user_balance/service/internal/logger"
)

type FileServer interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type fileServer struct {
	logger logger.Logger
}

func NewFileServer(logger logger.Logger) FileServer {
	return &fileServer{
		logger: logger,
	}
}

func (c *fileServer) Handle(w http.ResponseWriter, r *http.Request) {
	baseStoragePath := helpers.GetBaseStoragePath()

	httpFileServerHandler := http.FileServer(http.Dir(baseStoragePath))
	http.StripPrefix("/static", httpFileServerHandler).ServeHTTP(w, r)
}
