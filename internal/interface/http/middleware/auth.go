package middleware

import (
	"errors"
	"strings"

	"github.com/banggibima/go-fiber-jwt/config"
	"github.com/banggibima/go-fiber-jwt/internal/interface/http/presenter"
	"github.com/banggibima/go-fiber-jwt/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type AuthMiddleware struct {
	ResponsePresenter *presenter.ResponsePresenter
	Config            *config.Config
}

func NewAuthMiddleware(
	responsePresenter *presenter.ResponsePresenter,
	config *config.Config,
) *AuthMiddleware {
	return &AuthMiddleware{
		ResponsePresenter: responsePresenter,
		Config:            config,
	}
}

func (m *AuthMiddleware) Authentication(c *fiber.Ctx) error {
	token := c.Get("Authorization")
	if token == "" {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", errors.New("missing token"))
	}

	if !strings.HasPrefix(token, "Bearer ") {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", errors.New("invalid token"))
	}

	token = strings.TrimPrefix(token, "Bearer ")

	claims, err := jwt.ValidateToken(m.Config, token)
	if err != nil {
		return m.ResponsePresenter.SetError(c, fiber.StatusUnauthorized, "unauthorized", err)
	}

	c.Locals("claims", claims)

	return c.Next()
}
