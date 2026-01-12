package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Params struct {
	Skip  int64
	Limit int64
}

type Response struct {
	Skip       int64 `json:"skip"`
	Limit      int64 `json:"limit"`
	Total      int64 `json:"total"`
	TotalPages int64 `json:"total_pages"`
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

func NewResponse(params Params, total int64) Response {
	totalPages := total / params.Limit
	if total%params.Limit > 0 {
		totalPages++
	}

	return Response{
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
