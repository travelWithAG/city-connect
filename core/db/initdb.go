package db

import (
	"connect/core/migrations"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgreSQLDB() (*gorm.DB, error) {
	cfg := loadPsqlConfig()

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s timezone=UTC",
		cfg.DBHost,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBPort,
		cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to connect to PostgreSQL database: %v", err)
	}

	migrationError := migrations.Run(db)
	if migrationError != nil {
		return nil, fmt.Errorf("Failed to migrate database entities: %v", migrationError)
	}

	fmt.Println("Successfully connected to PostgreSQL database")
	
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25) // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(25) // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(5 * 60) // Maximum amount of time a connection may be reused (in seconds)

	log.Println("[DATABASE] PostgreSQL connection established successfully.")
	return db, nil

}

func InitMysqlDB() (*gorm.DB, error) {
	cfg := loadMySQLConfig()

	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to MySQL database: %v", err)
	}

	migrationError := migrations.Run(db)
	if migrationError != nil {
		return nil, fmt.Errorf("Failed to migrate database entities: %v", migrationError)
	}

	fmt.Print("Successfully connected to MySQL database;\n [MIGRATIONS] Has performed successifully;\n")

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("Failed to get database instance: %v", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxOpenConns(25) // Maximum number of open connections to the database
	sqlDB.SetMaxIdleConns(25) // Maximum number of idle connections in the pool
	sqlDB.SetConnMaxLifetime(5 * 60) // Maximum amount of time a connection may be reused (in seconds)

	log.Println("[DATABASE] MySQL connection established successfully.")
	return db, nil
}