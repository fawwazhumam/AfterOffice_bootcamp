package route

import (
	"go_crud/config"
	"go_crud/handler"
	"go_crud/middleware"

	"github.com/gofiber/fiber/v2"
)


func RouteInit(r *fiber.App) {

	r.Static("/public", config.ProjectRootPath + "/public/asset")
	//login
	r.Post("/login", handler.LoginHandler)

	r.Get("/user", middleware.Auth, handler.UserHandlerGetAll)
	r.Get("/user/:id", handler.UserHandlerGetById)
	r.Post("/user", handler.UserHandlerCreate)
	r.Put("/user/:id", handler.UserHandlerUpdate)
	r.Delete("/user/:id", handler.UserHandlerDelete)

	//shop
	r.Post("/shop", handler.ShopHandlerCreate)
	r.Get("/shop/:id", handler.ShopHandlerGetById)

}