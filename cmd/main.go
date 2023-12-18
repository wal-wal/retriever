package main

import (
	"application/domain/mantra/entity"
	"application/domain/mantra/use_case"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"persistence/persistence"
	"presentation/presentation"
	"presentation/router"
	"time"
)

func main() {
	app := fiber.New()
	t := time.Now()
	m := make(map[uuid.UUID]entity.Mantra)
	id := uuid.New()
	m[id] = entity.Mantra{Id: id, Writer: "김철우", Content: "내 코드는 레거시 코드다", CreatedAt: t.Add(-time.Minute)}

	mantraRepository := persistence.NewMantraPersistenceAdapter(m)
	readAllMantraUseCase := use_case.NewReadAllMantrasUseCase(mantraRepository)
	createMantraUseCase := use_case.NewCreateMantraUseCase(mantraRepository)
	deleteMantraUseCase := use_case.NewDeleteMantraUseCase(mantraRepository)
	mantraUseCase := use_case.NewMantraUseCase(*readAllMantraUseCase, *createMantraUseCase, *deleteMantraUseCase)
	mantraWebAdapter := presentation.New(mantraUseCase)
	mantraRouter := router.NewMantraRouter(*mantraWebAdapter)
	mantraRouter.Initialize(app)

	_ = app.Listen(":8080")
}
