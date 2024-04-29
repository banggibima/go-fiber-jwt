package http

import (
	"github.com/banggibima/go-fiber-jwt/internal/interface/http/handler"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	App                      *fiber.App
	UserHandler              *handler.UserHandler
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
