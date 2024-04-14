package routes

import (
	"banner_service/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler handlers.BannerHandler) {
	api := app.Group("/api/v1")

	api.Get("/user_banner", handler.GetUserBanner)
	api.Get("/banner", handler.GetBanner)
}
