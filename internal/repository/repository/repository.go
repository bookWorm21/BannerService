package repository

import (
	banner "banner_service/internal/model"
	"banner_service/internal/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ repository.GettingBannerRepository = (*GettingRepository)(nil)

type GettingRepository struct {
	postgres *pgxpool.Pool
}

func NewGettingRepository(postgres *pgxpool.Pool) *GettingRepository {
	return &GettingRepository{postgres: postgres}
}

func (repository *GettingRepository) Get(info banner.Info) (banner.Banner, error) {
	return banner.Banner{Content: `{"title": "some_title"}`}, nil
}
