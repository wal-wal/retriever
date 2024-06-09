package mantra_persentation

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"retriever-core/mantra/dto/request"
	"retriever-core/mantra/primary"
)

type MantraWebAdapter struct {
	mantraPort mantra_primary.MantraPrimaryPort
}

func NewMantraWebAdapter(mantraPort mantra_primary.MantraPrimaryPort) *MantraWebAdapter {
	return &MantraWebAdapter{
		mantraPort: mantraPort,
	}
}

func (r *MantraWebAdapter) ReadAllMantras(ctx *fiber.Ctx) error {
	mantras, err := r.mantraPort.ReadAllMantras()
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.JSON(mantras)
}

func (r *MantraWebAdapter) CreateMantra(ctx *fiber.Ctx) error {
	dto := new(mantra_request.CreateMantraReqDTO)
	_ = ctx.BodyParser(dto)
	if err := createMantraReqDTOValidator(*dto); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	if err := r.mantraPort.CreateMantra(*dto); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.SendStatus(201)
}

func (r *MantraWebAdapter) DeleteMantra(ctx *fiber.Ctx) error {
	mantraId := ctx.Params("mantraId")
	if err := deleteMantraQueryStringValidator(mantraId); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	err := r.mantraPort.DeleteMantra(uuid.MustParse(mantraId))
	if err != nil {
		return ctx.SendStatus(500)
	}
	return ctx.SendStatus(201)
}

func (r *MantraWebAdapter) ReadMantra(ctx *fiber.Ctx) error {
	mantraId := ctx.Params("mantraId")
	if err := readMantraQueryStringValidator(mantraId); err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	mantraResponse, err := r.mantraPort.ReadMantra(uuid.MustParse(mantraId))
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.Status(200).JSON(mantraResponse)
}

func (r *MantraWebAdapter) ReadMantraByRandom(ctx *fiber.Ctx) error {
	dto, err := r.mantraPort.ReadMantraByRandom()
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.Status(200).JSON(dto)
}

func (r *MantraWebAdapter) ReadMantrasBySpeaker(ctx *fiber.Ctx) error {
	speakerName := ctx.Query("speaker")
	mantras, err := r.mantraPort.ReadMantrasBySpeaker(speakerName)
	if err != nil {
		return ctx.Status(500).JSON(map[string]string{"message": err.Error()})
	}
	return ctx.Status(200).JSON(mantras)
}

func createMantraReqDTOValidator(dto mantra_request.CreateMantraReqDTO) error {
	if dto.Speaker == "" {
		return errors.New("발언자를 입력해주세요")
	}
	if dto.Content == "" {
		return errors.New("내용을 입력해주세요")
	}
	return nil
}

func readMantraQueryStringValidator(mantraId string) error {
	if mantraId == "" {
		return errors.New("아이디를 입력해주세요")
	}
	return nil
}

func deleteMantraQueryStringValidator(mantraId string) error {
	if mantraId == "" {
		return errors.New("아이디를 입력해주세요")
	}
	return nil
}
