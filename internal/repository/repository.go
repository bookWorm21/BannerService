package repository

import (
	banner "banner_service/internal/model"
)

type GettingBannerRepository interface {
	Get(info banner.Info) (banner.Banner, error)
}
