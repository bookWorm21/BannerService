package repository

import (
	banner "banner_service/internal/model"
	"banner_service/internal/repository"
	"context"
	"encoding/json"
	repErrors "errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

const (
	getBannerQuery = `
						SELECT banners.content
						FROM banner_schema.banners as banners
						LEFT JOIN banner_schema.banner_feature as banner_features 
						ON banners.id = banner_features.banner_id
						LEFT JOIN banner_schema.banner_tag as banner_tag
						ON banners.id = banner_tag.banner_id
						WHERE tag_id = $1 AND feature_id = $2 AND enable = $3`
)

var _ repository.GettingBannerRepository = (*GettingRepository)(nil)

type GettingRepository struct {
	postgres *pgxpool.Pool
}

func NewGettingRepository(postgres *pgxpool.Pool) *GettingRepository {
	return &GettingRepository{postgres: postgres}
}

func (repository *GettingRepository) Get(ctx *fiber.Ctx, info banner.Info) (banner.Banner, error) {
	if info.UseLastRevision {
		// TO:DO use cache
	}

	var content string
	var row = repository.postgres.QueryRow(context.Background(), getBannerQuery, info.TagId, info.FeatureId, true)
	err := row.Scan(&content)
	if err != nil {
		if repErrors.Is(err, pgx.ErrNoRows) {
			ctx.Status(fiber.StatusNotFound)
			return banner.Banner{}, errors.Wrapf(NotFound, "Banner feature - %d, tag - %d", info.FeatureId, info.TagId)
		}

		return banner.Banner{}, err
	}
	return banner.Banner{Content: json.RawMessage(content)}, nil
}
