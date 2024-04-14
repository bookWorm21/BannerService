package banner

import (
	banner "banner_service/internal/model"
	"banner_service/internal/repository"
	"banner_service/internal/service"
)

var _ service.BannerService = (*Service)(nil)

type Service struct {
	repository repository.GettingBannerRepository
}

func NewService(repository repository.GettingBannerRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetUserBanner(info banner.Info) (banner.Banner, error) {
	return s.repository.Get(info)
}
