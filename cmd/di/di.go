package di

import (
	"database/sql"
	"retriever-core/auth/use_case"
	"retriever-core/mantra/use_case"
	"retriever-core/user/use_case"
	"retriever-persistence/mantra"
	"retriever-persistence/mantra/repository"
	"retriever-persistence/user"
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
	mantraRepository := mantra_repository.NewMantraRepository(r.db)
	userRepository := user_repository.NewUserRepository(r.db)

	mantraPersistenceAdapter := mantra_persistence.NewMantraPersistenceAdapter(mantraRepository)
	userPersistenceAdapter := user_persistence.NewUserPersistenceAdapter(userRepository)

	readAllMantraUseCase := mantra_use_case.NewReadAllMantrasUseCase(mantraPersistenceAdapter)
	createMantraUseCase := mantra_use_case.NewCreateMantraUseCase(mantraPersistenceAdapter)
	deleteMantraUseCase := mantra_use_case.NewDeleteMantraUseCase(mantraPersistenceAdapter)
	readMantraUseCase := mantra_use_case.NewReadMantraUseCase(mantraPersistenceAdapter)
	mantraUseCase := mantra_use_case.NewMantraUseCase(readAllMantraUseCase,
		createMantraUseCase,
		deleteMantraUseCase,
		readMantraUseCase)

	createUserUseCase := user_use_case.NewCreateUserUseCase(userPersistenceAdapter)
	userUseCase := user_use_case.NewUserUseCase(*createUserUseCase)

	signInUseCase := auth_use_case.NewSignInUseCase(userPersistenceAdapter)
	authUseCase := auth_use_case.NewAuthUseCase(signInUseCase)

	userWebAdapter := user_presentation.NewUserWebAdapter(userUseCase)
	mantraWebAdapter := mantra_persentation.NewMantraWebAdapter(mantraUseCase)
	authWebAdapter := auth_presentation.NewAuthWebAdapter(authUseCase)

	return router.NewRouter(*mantraWebAdapter, *userWebAdapter, *authWebAdapter)
}
