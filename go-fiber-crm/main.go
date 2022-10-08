package main

import (
	"fmt"

	"github.com/youngshawn/11-go-projects/go-fiber-crm/database"
	"github.com/youngshawn/11-go-projects/go-fiber-crm/lead"

	"github.com/gofiber/fiber"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open(sqlite.Open("leads.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection opened to database.")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated.")
}

func main() {
	app := fiber.New()
	initDatabase()
	sqlDB, _ := database.DBConn.DB()
	defer sqlDB.Close()
	setupRoutes(app)
	app.Listen(3000)
}
