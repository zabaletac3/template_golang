package app

import (
	"github.com/gin-gonic/gin"

	"github.com/eren_dev/go_server/internal/modules/users"
	"github.com/eren_dev/go_server/internal/shared/database"
	"github.com/eren_dev/go_server/internal/shared/httpx"
)

func registerRoutes(engine *gin.Engine, db *database.MongoDB) {
	r := httpx.NewRouter(engine)

	api := r.Group("/api")

	if db != nil {
		users.RegisterRoutes(api, db)
	}
}
