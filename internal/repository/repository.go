package repository

import (
	banner "banner_service/internal/model"
	"github.com/gofiber/fiber/v2"
)

type GettingBannerRepository interface {
	Get(ctx *fiber.Ctx, info banner.Info) (banner.Banner, error)
}
