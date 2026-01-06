package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/eren_dev/go_server/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config) *Server {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	registerRoutes(router)

	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + cfg.Port,
			Handler:      router,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  120 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

