package users

import (
	"github.com/gin-gonic/gin"

	"github.com/eren_dev/go_server/internal/shared/pagination"
	"github.com/eren_dev/go_server/internal/shared/validation"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *gin.Context) (any, error) {
	var dto CreateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil, validation.Validate(err)
	}

	return h.service.Create(c.Request.Context(), &dto)
}

func (h *Handler) FindAll(c *gin.Context) (any, error) {
	params := pagination.FromContext(c)
	return h.service.FindAll(c.Request.Context(), params)
}

func (h *Handler) FindByID(c *gin.Context) (any, error) {
	id := c.Param("id")
	return h.service.FindByID(c.Request.Context(), id)
}

func (h *Handler) Update(c *gin.Context) (any, error) {
	id := c.Param("id")

	var dto UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil, validation.Validate(err)
	}

	return h.service.Update(c.Request.Context(), id, &dto)
}

func (h *Handler) Delete(c *gin.Context) (any, error) {
	id := c.Param("id")

	if err := h.service.Delete(c.Request.Context(), id); err != nil {
		return nil, err
	}

	return gin.H{"deleted": true}, nil
}
