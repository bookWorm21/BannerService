package banner

import (
	"banner_service/internal/api/handlers"
	"banner_service/internal/api/handlers/requests"
	banner "banner_service/internal/model"
	"banner_service/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
)

const (
	UseLastVersionDefault bool = false
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
	var request requests.UserBannerRequest
	err := ctx.QueryParser(&request)
	if err != nil {
		return err
	}

	var bannerInfo = banner.Info{TagId: request.TagId, FeatureId: request.FeatureId, UseLastVersion: UseLastVersionDefault}
	if request.UseLastVersion != nil {
		bannerInfo.UseLastVersion = *request.UseLastVersion
	}
	result, err := h.Service.GetUserBanner(bannerInfo)
	if err != nil {
		return err
	}

	err = ctx.JSON(&result)
	if err != nil {
		return err
	}
	return nil
}
