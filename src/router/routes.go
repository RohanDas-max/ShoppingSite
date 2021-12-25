package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rohandas-max/ambassador/src/controller"
	"github.com/rohandas-max/ambassador/src/middlewares"
)

func Setup(app *fiber.App) {

	api := app.Group("/api")

	//!! ADMIN routes
	admin := api.Group("/admin")
	admin.Post("/register", controller.Register)
	admin.Post("/login", controller.Login)

	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)

	adminAuthenticated.Post("/logout", controller.Logout)
	adminAuthenticated.Get("/user", controller.User)
	adminAuthenticated.Put("/users/info", controller.UpdateInfo)
	adminAuthenticated.Put("/users/updatepass", controller.UpdatePassword)
	adminAuthenticated.Get("/ambassador", controller.Ambassador)
	adminAuthenticated.Get("/products", controller.Products)
	adminAuthenticated.Post("/product", controller.CreateProduct)
	adminAuthenticated.Get("/product/:id", controller.GetProduct)
	adminAuthenticated.Put("/product/:id", controller.UpdateProduct)
	adminAuthenticated.Delete("/product/:id", controller.DeleteProduct)
	adminAuthenticated.Get("/users/:id/links", controller.Link)
	adminAuthenticated.Get("/orders", controller.Orders)

	ambassador := api.Group("/ambassador")
	ambassador.Post("/register", controller.Register)
	ambassador.Post("/login", controller.Login)
	ambassadorAuthenticated := ambassador.Use(middlewares.IsAuthenticated)
	ambassadorAuthenticated.Post("/logout", controller.Logout)
	ambassadorAuthenticated.Get("/user", controller.User)
	ambassadorAuthenticated.Put("/users/info", controller.UpdateInfo)
	ambassadorAuthenticated.Put("/users/updatepass", controller.UpdatePassword)
	ambassadorAuthenticated.Get("/products/fe", controller.ProductsFE)
	ambassadorAuthenticated.Get("/products/be", controller.ProductsBE)
}
