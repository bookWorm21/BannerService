package service

import "banner_service/internal/model/banner"

type BannerService interface {
	GetUserBanner(request banner.UserBannerRequest) (banner.UserBannerResponse, error)
}
