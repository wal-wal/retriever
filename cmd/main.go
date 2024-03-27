package main

import (
	"application/domain/mantra/entity"
	"application/domain/mantra/use_case"
	useCase2 "application/domain/user/use_case"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"os"
	repository2 "persistence/mantra/repository"
	"persistence/user/repository"
	"presentation/mantra"
	presentation2 "presentation/user"
	"time"
)

func main() {
	app := fiber.New()
	t := time.Now()
	password := os.Getenv("PASSWORD")
	db, err := sql.Open("mysql", "root:"+password+"@tcp(localhost:3306)/WALWAL")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	m := make(map[uuid.UUID]entity.Mantra)
	id := uuid.New()
	m[id] = entity.Mantra{Id: id, Writer: "김철우", Content: "내 코드는 레거시 코드다", CreatedAt: t.Add(-time.Minute)}
	mantraRepository := repository2.NewMantraPersistenceAdapter(m)
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
