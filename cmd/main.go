package main

import (
	"application/domain/mantra/use_case"
	useCase2 "application/domain/user/use_case"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"os"
	repository2 "persistence/mantra/repository"
	"persistence/user/repository"
	"presentation/mantra"
	presentation2 "presentation/user"
)

func main() {
	app := fiber.New()
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/WALWAL")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	mantraRepository := repository2.NewMantraPersistenceAdapter(db)
	readAllMantraUseCase := use_case.NewReadAllMantrasUseCase(mantraRepository)
	createMantraUseCase := use_case.NewCreateMantraUseCase(mantraRepository)
	deleteMantraUseCase := use_case.NewDeleteMantraUseCase(mantraRepository)
	mantraUseCase := use_case.NewMantraUseCase(*readAllMantraUseCase, *createMantraUseCase, *deleteMantraUseCase)
	mantraWebAdapter := mantra.New(mantraUseCase)
	mantraRouter := mantra.NewMantraRouter(*mantraWebAdapter)
	mantraRouter.Initialize(app)

	userRepository := repository.NewUserPersistenceAdapter(db)
	createUserUseCase := useCase2.NewCreateUserUseCase(userRepository)
	userUseCase := useCase2.NewUserUseCase(*createUserUseCase)
	userWebAdapter := presentation2.NewUserWebAdapter(userUseCase)
	userRouter := presentation2.NewUserRouter(*userWebAdapter)
	userRouter.Initialize(app)
	_ = app.Listen(":8080")
}
