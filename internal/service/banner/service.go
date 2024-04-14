package banner

import (
	banner "banner_service/internal/model"
	"banner_service/internal/repository"
	"banner_service/internal/service"
	"github.com/gofiber/fiber/v2"
)

var _ service.BannerService = (*Service)(nil)

type Service struct {
	repository repository.GettingBannerRepository
}

func NewService(repository repository.GettingBannerRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUserBanner(ctx *fiber.Ctx, info banner.Info) (banner.Banner, error) {
	return s.repository.Get(ctx, info)
}
