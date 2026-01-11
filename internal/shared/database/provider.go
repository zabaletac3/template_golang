package database

import "github.com/eren_dev/go_server/internal/config"

func NewProvider(cfg *config.Config) (*MongoDB, error) {
	if cfg.MongoDatabase == "" {
		return nil, nil
	}
	return NewMongoDB(cfg)
}
