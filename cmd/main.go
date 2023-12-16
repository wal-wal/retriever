package main

import (
	"application/domain/mantra/dto/request"
	"application/domain/mantra/entity"
	"application/domain/mantra/use_case"
	"github.com/gofiber/fiber/v2"
	"persistence/persistence"
	"presentation/presentation"
	"time"
)

func main() {
	app := fiber.New()
	t := time.Now()
	m := make(map[int]entity.Mantra)
	m[0] = entity.Mantra{Id: 1, Writer: "김철우", Content: "내 코드는 레거시 코드다", CreatedAt: t.Add(-time.Minute)}
	m[1] = entity.Mantra{Id: 2, Writer: "정지관", Content: "내가 짱이다. 남에게 주눅들 이유 따위는 없다", CreatedAt: t.Add(-time.Minute * 2)}
	m[2] = entity.Mantra{Id: 3, Writer: "이정호", Content: "푸른 언덕에~~", CreatedAt: t.Add(-time.Minute * 3)}
	m[3] = entity.Mantra{Id: 4, Writer: "손영달", Content: "엄마한테서 최대한 3억을 뜯어내라", CreatedAt: t.Add(-time.Minute * 4)}
	mantraRepository := persistence.New(m)
	readAllMantraUseCase := use_case.NewReadAllMantrasUseCase(mantraRepository)
	createMantraUseCase := use_case.NewCreateMantraUseCase(mantraRepository)
	mantraController := presentation.New(*readAllMantraUseCase, *createMantraUseCase)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(mantraController.ReadAllMantras())
	})

	app.Post("/", func(ctx *fiber.Ctx) error {
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
