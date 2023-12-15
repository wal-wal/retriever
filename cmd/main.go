package main

import (
	"github.com/gofiber/fiber/v2"
	"retriever/application/domain/mantra/entity"
	uc "retriever/application/domain/mantra/use_case"
	ctrl "retriever/infrastructure/primary_adapter/presentation"
	repo "retriever/infrastructure/secondary_adapter/persistence"
)

func main() {
	app := fiber.New()

	m := make(map[int]entity.Mantra)
	m[0] = entity.Mantra{Id: 1, Writer: "김철우", Content: "내 코드는 레거시 코드다"}
	m[1] = entity.Mantra{Id: 2, Writer: "정지관", Content: "무력한 내가 그들에게 복수할 수 있는 방법은 성공하는 것 뿐이다"}
	m[2] = entity.Mantra{Id: 3, Writer: "이정호", Content: "푸른 언덕에~~"}
	m[3] = entity.Mantra{Id: 4, Writer: "손영달", Content: "엄마한테서 최대한 3억을 뜯어내라"}
	mantraRepository := repo.New(m)
	readAllMantraUseCase := uc.New(mantraRepository)
	mantraController := ctrl.New(*readAllMantraUseCase)

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(mantraController.ReadAllMantras())
	})
	_ = app.Listen(":8080")
}
