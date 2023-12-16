package main

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/entity"
	"application/domain/mantra/use_case"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"persistence/persistence"
	"presentation/presentation"
	"time"
)

func main() {
	app := fiber.New()
	t := time.Now()
	m := make(map[uuid.UUID]entity.Mantra)
	id := uuid.New()
	m[id] = entity.Mantra{Id: id, Writer: "김철우", Content: "내 코드는 레거시 코드다", CreatedAt: t.Add(-time.Minute)}
	mantraRepository := persistence.New(m)
	readAllMantraUseCase := use_case.NewReadAllMantrasUseCase(mantraRepository)
	createMantraUseCase := use_case.NewCreateMantraUseCase(mantraRepository)
	mantraController := presentation.New(*readAllMantraUseCase, *createMantraUseCase)

	mantra := app.Group("/mantra")
	mantra.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(mantraController.ReadAllMantras())
	})
	mantra.Post("/", func(ctx *fiber.Ctx) error {
		dto := new(request.CreateMantraReqDTO) // 나중엔 이거도 분리하기
		_ = ctx.BodyParser(dto)
		err := mantraController.CreateMantra(*dto)
		if err != nil { // 커스텀 에러 만들기
			return ctx.SendStatus(500)
		}
		return ctx.SendStatus(201)
	})
	_ = app.Listen(":8080")
}
