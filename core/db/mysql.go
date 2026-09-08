package db

import "connect/core/config"

func loadMySQLConfig() config.DbConfig {
	return config.LoadAppConfig("mysql").DbConfig
}