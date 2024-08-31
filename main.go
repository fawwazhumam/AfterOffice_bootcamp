package main

import (
	"go_crud/database"
	"go_crud/database/migration"
	"go_crud/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Inisial Database
	database.DatabaseInit()
	//Run Migration
	migration.RunMigration()

	app := fiber.New()

	//Inisial Route
	route.RouteInit(app)

	app.Listen(":3000")
}