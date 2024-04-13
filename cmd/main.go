package main

import (
	"banner_service/internal/api/handlers/banner"
	route_v1 "banner_service/internal/api/v1/routes/banner"
	"banner_service/internal/config"
	bannerService "banner_service/internal/service/banner"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	appConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()
	var bannerHandler = banner.NewHandler(bannerService.NewService())
	route_v1.SetupRoutes(app, bannerHandler)
	log.Fatal(app.Listen(appConfig.ServicePort))
}
