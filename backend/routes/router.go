package routes

import (
	"mangojek-backend/controller"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, userController controller.UserController, messageController controller.MessageController) {
	app.Get("/api/users", userController.FindAll)
	app.Post("/api/register", userController.Register)
	app.Post("/api/login", userController.Login)
	app.Post("/api/message", messageController.SendMessage)
	app.Get("/api/messages", messageController.GetMessages)
	// app.Post("/api/users", middleware.AuthMiddleware, controller.Insert)
}
