package banner

import (
	"banner_service/internal/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler handlers.BannerHandler) {
	// api := app.Group("/api/v1")

	app.Get("/user_banner", handler.GetUserBanner)
	app.Get("/banner", handler.GetBanner)
}
