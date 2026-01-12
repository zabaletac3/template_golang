package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Skip  int64
	Limit int64
}

// PaginationInfo información de paginación
// @name PaginationInfo
type PaginationInfo struct {
	// Registros saltados
	Skip int64 `json:"skip" example:"0"`
	// Límite de registros por página
	Limit int64 `json:"limit" example:"10"`
	// Total de registros
	Total int64 `json:"total" example:"100"`
	// Total de páginas
	TotalPages int64 `json:"total_pages" example:"10"`
}

const (
	DefaultLimit = 10
	MaxLimit     = 100
)

func FromContext(c *gin.Context) Params {
	skip := parseQueryInt64(c, "skip", 0)
	limit := parseQueryInt64(c, "limit", DefaultLimit)

	if skip < 0 {
		skip = 0
	}
	if limit <= 0 {
		limit = DefaultLimit
	}
	if limit > MaxLimit {
		limit = MaxLimit
	}

	return Params{
		Skip:  skip,
		Limit: limit,
	}
}

func NewPaginationInfo(params Params, total int64) PaginationInfo {
	totalPages := total / params.Limit
	if total%params.Limit > 0 {
		totalPages++
	}

	return PaginationInfo{
		Skip:       params.Skip,
		Limit:      params.Limit,
		Total:      total,
		TotalPages: totalPages,
	}
}

func parseQueryInt64(c *gin.Context, key string, def int64) int64 {
	val := c.Query(key)
	if val == "" {
		return def
	}
	n, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return def
	}
	return n
}
