package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type DbConfig struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBAutoMigrate bool
	DBShowSQL     bool
}

type AppConfig struct {
	AppName    string
	AppVersion string
	AppEnv     string
	AppPort    string
	DbConfig   DbConfig
	JWTSecret  string
	JWTExpiry  int
	RedisHost  string
	RedisPort  int
	RedisPassword string
	RedisDB    int
}



func LoadAppConfig(drv string) *AppConfig {
	// 2. Extract values using os.Getenv or fall back to defaults
	autoMigrate, _ := strconv.ParseBool(getEnv("DB_AUTO_MIGRATE", "true"))
	showSQL, _ := strconv.ParseBool(getEnv("DB_SHOW_SQL", "true"))

	if drv == ""{
		drv = "postgres"
	}

	cfg := &AppConfig{
		AppName:    getEnv("APP_NAME", "CityConnect"),
		AppVersion: getEnv("APP_VERSION", "latest"),
		AppEnv:     getEnv("APP_ENV", "development"),
		AppPort:    getEnv("APP_PORT", "8080"),
		DbConfig: DbConfig{
			DBDriver:   getEnv("DB_DRIVER", drv),
			DBHost:     getEnv("DB_HOST", "localhost"),
			DBPort:     getEnv("DB_PORT", "5432"),
			DBUser:     getEnv("DB_USER", "gouser"),
			DBPassword: getEnv("DB_PASSWORD", "g0lang2026"),
			DBName:     getEnv("DB_NAME", "connect_db"),
			DBSSLMode:  getEnv("DB_SSL_MODE", "disable"),
			DBAutoMigrate: autoMigrate,
			DBShowSQL:     showSQL,
		},
		JWTSecret:  getEnv("JWT_SECRET", "secret_88_jwt_key"),
		JWTExpiry:  getEnvAsInt(getEnv("JWT_EXPIRY", "24"), 24), // default to 24 hours
		RedisHost:  getEnv("REDIS_HOST", "localhost"),
		RedisPort:  getEnvAsInt("REDIS_PORT", 6379),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:    getEnvAsInt("REDIS_DB", 0),
	}

	fmt.Printf("[%s] Configuration Loaded for City Connect Platform, Running at [%s]:[%s]\n Database cfg [%s %s %s]", cfg.AppName, "127.0.0.1", cfg.AppPort, cfg.DbConfig.DBHost, cfg.DbConfig.DBPort, cfg.DbConfig.DBName)
	
	return cfg
}

func getEnv(key string, defaultValue string) string {
	found, exists := os.LookupEnv(key)
	fmt.Printf("Loading environment variable: %s\n Default value: %v Existing value is: %v", key, found, exists)
	if exists {
		return found
	}
	return defaultValue
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")

	if(valueStr == "") {
		return defaultVal
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Printf("Warning: Invalid integer value for %s. Using default %d", name, defaultVal)
		return defaultVal
	}

	return value
}
