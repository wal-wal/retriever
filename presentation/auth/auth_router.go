package auth

import (
	"github.com/gofiber/fiber/v2"
)

type AuthRouter struct {
	adapter AuthWebAdapter
}

func NewAuthRouter(adapter AuthWebAdapter) *AuthRouter {
	return &AuthRouter{
		adapter: adapter,
	}
}

func (r *AuthRouter) Initialize(app *fiber.App) {
	authRouter := app.Group("/auth")

	authRouter.Post("/sign-in", r.adapter.SignIn)
}
