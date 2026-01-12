package auth

import (
	"github.com/eren_dev/go_server/internal/config"
	"github.com/eren_dev/go_server/internal/modules/users"
	"github.com/eren_dev/go_server/internal/shared/auth"
	"github.com/eren_dev/go_server/internal/shared/database"
	"github.com/eren_dev/go_server/internal/shared/httpx"
)

func RegisterRoutes(public *httpx.Router, private *httpx.Router, db *database.MongoDB, cfg *config.Config) {
	userRepo := users.NewRepository(db)
	jwtService := auth.NewJWTService(cfg)
	service := NewService(userRepo, jwtService)
	handler := NewHandler(service)

	// Public routes
	public.POST("/auth/register", handler.Register)
	public.POST("/auth/login", handler.Login)
	public.POST("/auth/refresh", handler.Refresh)

	// Protected routes
	private.GET("/auth/me", handler.Me)
}
