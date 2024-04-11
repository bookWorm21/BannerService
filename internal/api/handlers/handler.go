package handlers

import "github.com/gofiber/fiber/v2"

type BannerHandler interface {
	GetUserBanner(ctx *fiber.Ctx) error
	GetBanner(ctx *fiber.Ctx) error
}
