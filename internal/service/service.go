package service

import (
	"banner_service/internal/model"
	"github.com/gofiber/fiber/v2"
)

type BannerService interface {
	GetUserBanner(ctx *fiber.Ctx, info banner.Info) (banner.Banner, error)
}
