package mantra

import (
	"github.com/gofiber/fiber/v2"
)

type MantraRouter struct {
	adapter MantraWebAdapter
}

func NewMantraRouter(adapter MantraWebAdapter) *MantraRouter {
	return &MantraRouter{
		adapter: adapter,
	}
}

func (r *MantraRouter) Initialize(app *fiber.App) {
	mantraRouter := app.Group("/mantra")

	mantraRouter.Get("/", r.adapter.ReadAllMantras)
	mantraRouter.Post("/", r.adapter.CreateMantra)
	mantraRouter.Delete("/:id", r.adapter.DeleteMantra)
}
