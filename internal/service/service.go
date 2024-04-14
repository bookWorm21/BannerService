package service

import (
	"banner_service/internal/model"
)

type BannerService interface {
	GetUserBanner(info banner.Info) (banner.Banner, error)
}
