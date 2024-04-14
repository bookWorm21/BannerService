package routes

import (
	"banner_service/external/auth"
	"banner_service/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRoutes(app *fiber.App, handler handlers.BannerHandler) {
	api := app.Group("/api/v1", logger.New())

	api.Get("/user_banner", auth.HaveAccessToUser, handler.GetUserBanner)
	api.Get("/banner", auth.HaveAccessToAdmin, handler.GetBanner)
}
