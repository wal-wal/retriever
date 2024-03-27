package presentation

import "github.com/gofiber/fiber/v2"

type UserRouter struct {
	adapter UserWebAdapter
}

func NewUserRouter(adapter UserWebAdapter) *UserRouter {
	return &UserRouter{
		adapter: adapter,
	}
}

func (r *UserRouter) Initialize(app *fiber.App) {
	userRouter := app.Group("/user")

	userRouter.Post("/sign-up", r.adapter.CreateUser)
}
