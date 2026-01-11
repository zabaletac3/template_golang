package health

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/eren_dev/go_server/internal/shared/database"
)

var db *database.MongoDB

func SetDatabase(d *database.MongoDB) {
	db = d
}

func Health(c *gin.Context) {
	status := gin.H{"status": "ok"}

	if db != nil {
		if err := db.Health(c.Request.Context()); err != nil {
			status["database"] = "unhealthy"
			c.JSON(http.StatusServiceUnavailable, status)
			return
		}
		status["database"] = "healthy"
	}

	c.JSON(http.StatusOK, status)
}

func Ready(c *gin.Context) {
	if !IsReady() {
		c.JSON(http.StatusServiceUnavailable, gin.H{"status": "not_ready"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
