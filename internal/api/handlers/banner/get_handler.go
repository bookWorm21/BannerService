package banner

import (
	"banner_service/internal/api/handlers"
	banner "banner_service/internal/model"
	"banner_service/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"log"
)

const (
	UseLastRevisionDefault bool = false
)

const (
	TagId           string = "tag_id"
	FeatureId       string = "feature_id"
	UseLastRevision string = "use_last_revision"
)

var _ handlers.BannerHandler = (*Handler)(nil)

type Handler struct {
	Service service.BannerService
}

func NewHandler(service service.BannerService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetBanner(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (h *Handler) GetUserBanner(ctx *fiber.Ctx) error {
	var headers = ctx.GetReqHeaders()
	tokenValue, inMap := headers["Token"]
	if inMap && len(tokenValue) > 0 {
		log.Print(tokenValue[0])
	}

	var tagId = ctx.QueryInt(TagId, -1)
	if tagId == -1 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.Wrapf(NotCorrectData, "Tag id")
	}

	var featureId = ctx.QueryInt(FeatureId, -1)
	if featureId == -1 {
		ctx.Status(fiber.StatusBadRequest)
		return errors.Wrapf(NotCorrectData, "Feature id")
	}

	var bannerInfo = banner.Info{TagId: tagId, FeatureId: featureId, UseLastRevision: UseLastRevisionDefault}
	bannerInfo.UseLastRevision = ctx.QueryBool(UseLastRevision, UseLastRevisionDefault)

	result, err := h.Service.GetUserBanner(ctx, bannerInfo)
	if err != nil {
		return err
	}

	err = ctx.Status(200).JSON(result.Content)
	if err != nil {
		return err
	}
	return nil
}
