package controllers

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"notion/internal/interfaces"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type HackersController struct {
	Log          *zap.SugaredLogger
	RedisService interfaces.IRedisService
	Validator    *validator.Validate
}

func (c *HackersController) GetHackers(fctx *fiber.Ctx) error {
	// можно добавить какой-нибудь message по типу no results и соответствующий код, но в тз этого не указано по-этому так :)
	return fctx.Status(http.StatusOK).JSON(c.RedisService.GetHackers())
}
