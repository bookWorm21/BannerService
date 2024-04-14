package main

import (
	"banner_service/internal/api/handlers/banner"
	route_v1 "banner_service/internal/api/v1/routes"
	"banner_service/internal/config"
	"banner_service/internal/repository/repository"
	service "banner_service/internal/service/banner"
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func connectToPostgres(appConfig *config.Config) *pgxpool.Pool {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		appConfig.PostgresUser,
		appConfig.PostgresPass,
		appConfig.PostgresHost,
		appConfig.PostgresPort,
		appConfig.PostgresName)

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(fmt.Errorf("unable to connect to database: %v", err))
	}

	if err = pool.Ping(context.Background()); err != nil {
		panic(fmt.Errorf("failed ping to postgres: %v", err))
	}

	log.Println("Successfully connected to postgres")
	return pool
}

func main() {
	appConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	postgres := connectToPostgres(&appConfig)
	defer postgres.Close()

	bannerRepository := repository.NewGettingRepository(postgres)
	bannerService := service.NewService(bannerRepository)
	var bannerHandler = banner.NewHandler(bannerService)

	app := fiber.New()
	route_v1.SetupRoutes(app, bannerHandler)
	log.Fatal(app.Listen(appConfig.ServicePort))
}
