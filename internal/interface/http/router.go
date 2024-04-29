package http

import (
	"github.com/banggibima/go-fiber-jwt/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App                      *fiber.App
	UserHandler              *handler.UserHandler
	TokenHandler             *handler.TokenHandler
	AuthenticationMiddleware fiber.Handler
}

func NewRouter(
	app *fiber.App,
) *Router {
	return &Router{
		App: app,
	}
}

func (r *Router) Public() {
	api := r.App.Group("/api")

	auth := api.Group("/auth")
	auth.Post("/login", r.UserHandler.Login)
	auth.Post("/register", r.UserHandler.Register)

	token := api.Group("/token")
	token.Get("/:refresh_token", r.TokenHandler.ReadByRefreshToken)
	token.Post("/", r.TokenHandler.Create)
	token.Delete("/:refresh_token", r.TokenHandler.DeleteByRefreshToken)
}

func (r *Router) Protected() {
	api := r.App.Group("/api")

	user := api.Group("/users", r.AuthenticationMiddleware)
	user.Get("/", r.UserHandler.ReadAll)
	user.Get("/:id", r.UserHandler.ReadByID)
	user.Get("/username/:username", r.UserHandler.ReadByUsername)
	user.Post("/", r.UserHandler.Create)
	user.Put("/:id", r.UserHandler.Update)
	user.Delete("/:id", r.UserHandler.Delete)
}
