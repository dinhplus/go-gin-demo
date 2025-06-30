package repository

import (
	"fmt"
	"log"
	"my-gin-app/internal/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, relying on system environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PGHOST"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGPORT"),
	)
	fmt.Println("Connecting to database with DSN:", dsn)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB.AutoMigrate(
		&model.User{},
		&model.Department{},
		&model.Stack{},
		&model.Position{},
		&model.Role{},
		&model.Permission{},
	)
	SeedRoles()
	SeedDepartments()
	SeedStacks()
	SeedPositions()
}
