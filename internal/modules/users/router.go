package users

import (
	"github.com/eren_dev/go_server/internal/shared/database"
	"github.com/eren_dev/go_server/internal/shared/httpx"
)

func RegisterRoutes(r *httpx.Router, db *database.MongoDB) {
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	users := r.Group("/users")

	users.POST("", handler.Create)
	users.GET("", handler.FindAll)
	users.GET("/:id", handler.FindByID)
	users.PUT("/:id", handler.Update)
	users.DELETE("/:id", handler.Delete)
}
