package banner

import (
	"banner_service/internal/api/handlers"
	models "banner_service/internal/model/banner"
	"banner_service/internal/service"
	"github.com/gofiber/fiber/v2"
	"log"
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
	var request models.UserBannerRequest
	err := ctx.QueryParser(&request)
	if err != nil {
		return err
	}

	result, err := h.Service.GetUserBanner(request)
	if err != nil {
		return err
	}

	err = ctx.JSON(&result)
	if err != nil {
		return err
	}
	return nil
}
