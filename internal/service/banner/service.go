package banner

import (
	"banner_service/internal/model/banner"
	service2 "banner_service/internal/service"
	"log"
)

var _ service2.BannerService = (*Service)(nil)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (service *Service) GetUserBanner(request banner.UserBannerRequest) (banner.UserBannerResponse, error) {
	log.Println("GetUserBanner")
	return banner.UserBannerResponse{}, nil
}
