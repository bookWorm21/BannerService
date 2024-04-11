package main

import (
	"banner_service/internal/api/handlers/banner"
	route_v1 "banner_service/internal/api/v1/routes/banner"
	bannerService "banner_service/internal/service/banner"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	var bannerService = bannerService.NewService()
	var bannerHandler = banner.NewHandler(bannerService)
	route_v1.SetupRoutes(app, bannerHandler)
	log.Fatal(app.Listen(":3000"))
}
