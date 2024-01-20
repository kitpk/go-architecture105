package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/kitpk/go-architecture105/adapters"
	"github.com/kitpk/go-architecture105/core"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	app := fiber.New()

	// Initialize the database connection
	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&core.Order{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderServicer(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	app.Post("/order", orderHandler.CreateOrder)

	app.Listen(":8080")
}
