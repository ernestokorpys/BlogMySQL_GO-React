package routes

import (
	"github.com/ernestokorpys/gobackend/controller"
	"github.com/ernestokorpys/gobackend/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticate)
	app.Post("api/post", controller.CreatePost)
	app.Get("api/allpost", controller.AllPost)
	app.Get("api/allpost/:id", controller.DetailPost)
	app.Put("api/updatepost/:id", controller.UpdatePost)
	app.Get("api/uniquepost", controller.UniquePost)
	app.Delete("api/deletepost/:id", controller.DeletePost)
}
