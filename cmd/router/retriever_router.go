package router

import (
	"github.com/gofiber/fiber/v2"
	"retriever-presentation/auth"
	"retriever-presentation/mantra"
	"retriever-presentation/user"
)

type RetrieverRouter struct {
	mantra mantra_persentation.MantraWebAdapter
	user   user_presentation.UserWebAdapter
	auth   auth_presentation.AuthWebAdapter
}

func NewRouter(mantra mantra_persentation.MantraWebAdapter,
	user user_presentation.UserWebAdapter,
	auth auth_presentation.AuthWebAdapter) *RetrieverRouter {
	return &RetrieverRouter{
		mantra: mantra,
		user:   user,
		auth:   auth,
	}
}

func (r *RetrieverRouter) Initialize(app *fiber.App) {
	userRouter := app.Group("/user")
	mantraRouter := app.Group("/mantra")
	authRouter := app.Group("/auth")

	userRouter.Post("/sign-up", r.user.CreateUser)

	mantraRouter.Get("/:mantraId", r.mantra.ReadMantra)
	mantraRouter.Get("/", r.mantra.ReadAllMantras)
	mantraRouter.Post("/", r.mantra.CreateMantra)
	mantraRouter.Delete("/:id", r.mantra.DeleteMantra)

	authRouter.Post("/sign-in", r.auth.SignIn)
}
