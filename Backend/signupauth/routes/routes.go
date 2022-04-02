package routes

import (
	"CAW/Backend/signupauth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register",controllers.Register)
	app.Post("/api/login",controllers.Login)
	app.Get("/api/user",controllers.User)
	app.Post("/api/logout",controllers.Logout)
	app.Post("/api/booking",controllers.Booking)
	app.Get("/api/orders",controllers.GetOrders)
	app.Get("/api/orders/id",controllers.GetOrderById)
	app.Delete("/api/cancelluggage/id",controllers.CancelLuggageOrder)
}