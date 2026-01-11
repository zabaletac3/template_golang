package users

import (
	"github.com/gin-gonic/gin"
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
		return nil, err
	}

	return h.service.Create(c.Request.Context(), &dto)
}

func (h *Handler) FindAll(c *gin.Context) (any, error) {
	return h.service.FindAll(c.Request.Context())
}

func (h *Handler) FindByID(c *gin.Context) (any, error) {
	id := c.Param("id")
	return h.service.FindByID(c.Request.Context(), id)
}

func (h *Handler) Update(c *gin.Context) (any, error) {
	id := c.Param("id")

	var dto UpdateUserDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		return nil, err
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
