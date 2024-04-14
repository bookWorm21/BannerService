package auth

import (
	"github.com/gofiber/fiber/v2"
)

const (
	userToken  = "user_token"
	adminToken = "admin_token"
)

func HaveAccessToAdmin(ctx *fiber.Ctx) error {
	token, err := GetToken(ctx)
	if err != nil {
		return err
	}

	if token == adminToken {
		return nil
	}

	ctx.Status(403)
	return NoAccess
}

func HaveAccessToUser(ctx *fiber.Ctx) error {
	token, err := GetToken(ctx)
	if err != nil {
		return err
	}

	if token == userToken {
		return nil
	}

	ctx.Status(403)
	return NoAccess
}

func GetToken(ctx *fiber.Ctx) (string, error) {
	var headers = ctx.GetReqHeaders()
	tokenValue, inMap := headers["Token"]
	if inMap && len(tokenValue) == 0 {
		return tokenValue[0], nil
	}

	ctx.Status(fiber.StatusUnauthorized)
	return "", Unauthorized
}
