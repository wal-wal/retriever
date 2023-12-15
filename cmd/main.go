package main

import (
	"application/domain/mantra/entity"
	"application/domain/mantra/use_case"
	"github.com/gofiber/fiber/v2"
	"persistence/persistence"
	"presentation/presentation"
)

func main() {
	app := fiber.New()

	m := make(map[int]entity.Mantra)
	m[0] = entity.Mantra{Id: 1, Writer: "김철우", Content: "내 코드는 레거시 코드다"}
	m[1] = entity.Mantra{Id: 2, Writer: "정지관", Content: "내가 짱이다. 남에게 주눅들 이유 따위는 없다"}
	m[2] = entity.Mantra{Id: 3, Writer: "이정호", Content: "푸른 언덕에~~"}
	m[3] = entity.Mantra{Id: 4, Writer: "손영달", Content: "엄마한테서 최대한 3억을 뜯어내라"}
	mantraRepository := persistence.New(m)
	readAllMantraUseCase := use_case.New(mantraRepository)
	mantraController := presentation.New(*readAllMantraUseCase)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(mantraController.ReadAllMantras())
	})
	_ = app.Listen(":8080")
}
