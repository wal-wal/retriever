package di

import (
	"database/sql"
	"retriever-core/auth/use_case"
	"retriever-core/mantra/use_case"
	"retriever-core/user/use_case"
	"retriever-persistence/mantra/repository"
	"retriever-persistence/user/repository"
	"retriever-presentation/auth"
	"retriever-presentation/mantra"
	"retriever-presentation/user"
	"retriever/cmd/router"
)

type DependencyInjector struct {
	db *sql.DB
}

func NewDependencyInjector(db *sql.DB) *DependencyInjector {
	return &DependencyInjector{
		db: db,
	}
}

func (r *DependencyInjector) PureDI() *router.RetrieverRouter {
	mantraRepository := mantra_repository.NewMantraPersistenceAdapter(r.db)
	userRepository := user_repository.NewUserPersistenceAdapter(r.db)

	readAllMantraUseCase := mantra_use_case.NewReadAllMantrasUseCase(mantraRepository)
	createMantraUseCase := mantra_use_case.NewCreateMantraUseCase(mantraRepository)
	deleteMantraUseCase := mantra_use_case.NewDeleteMantraUseCase(mantraRepository)
	mantraUseCase := mantra_use_case.NewMantraUseCase(readAllMantraUseCase, createMantraUseCase, deleteMantraUseCase)

	createUserUseCase := user_use_case.NewCreateUserUseCase(userRepository)
	userUseCase := user_use_case.NewUserUseCase(*createUserUseCase)

	signInUseCase := auth_use_case.NewSignInUseCase(userRepository)
	authUseCase := auth_use_case.NewAuthUseCase(signInUseCase)

	userWebAdapter := user_presentation.NewUserWebAdapter(userUseCase)
	mantraWebAdapter := mantra_persentation.NewMantraWebAdapter(mantraUseCase)
	authWebAdapter := auth_presentation.NewAuthWebAdapter(authUseCase)
	return router.NewRouter(*mantraWebAdapter, *userWebAdapter, *authWebAdapter)
}
