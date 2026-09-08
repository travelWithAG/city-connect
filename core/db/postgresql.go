package db

import "connect/core/config"

func loadPsqlConfig() config.DbConfig {
	cfg := config.LoadAppConfig("postgres")
	return cfg.DbConfig
}