package routes

import (
	"backend/domains"
	"backend/services"
	"github.com/gofiber/fiber"
	"strconv"
)

func ConfigVagaRoute(service *services.VagaService, app *fiber.App) {
	vagasGroup := app.Group("/vagas")

	vagasGroup.Get("/", func(ctx *fiber.Ctx) {
		vaga, err := service.GetAll()
		if err != nil {
			ctx.Next(err)
			return
		}
		var vagaDTO []*domains.VagaSimplesDTO
		for _, v := range vaga {
			vagaDTO = append(vagaDTO, v.GetVagaSimplesDTO())
		}
		err = ctx.JSON(vagaDTO)
		if err != nil {
			ctx.Next(err)
		}
	})

	vagasGroup.Get("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		vaga, err := service.GetOne(id)
		if err != nil {
			ctx.Next(err)
			return
		}
		vagaDTO := vaga.GetVagaComplexaDTO()
		err = ctx.JSON(vagaDTO)
		if err != nil {
			ctx.Next(err)
		}
	})

	vagasGroup.Post("/", func(ctx *fiber.Ctx) {
		vagaDTO := &domains.VagaSimplesDTO{}
		if err := ctx.BodyParser(vagaDTO); err != nil {
			ctx.Next(err)
			return
		}

		vaga := vagaDTO.FromDTOSimples()
		vaga, err := service.Create(vaga)
		if err != nil {
			ctx.Next(err)
		}
		ctx.Set("Location", ctx.Path()+strconv.Itoa(int(vaga.ID)))
		ctx.Status(201)
		ctx.Send()
	})

	vagasGroup.Put("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		vagaDTO := &domains.VagaSimplesDTO{}
		if err := ctx.BodyParser(vagaDTO); err != nil {
			ctx.Next(err)
			return
		}

		vaga := vagaDTO.FromDTOSimples()
		vaga, err := service.Update(vaga, id)
		if err != nil {
			ctx.Next(err)
			return
		}

		err = ctx.JSON(vaga)
		if err != nil {
			ctx.Next(err)
		}
	})

	vagasGroup.Delete("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		err := service.Delete(id)
		if err != nil {
			ctx.Next(err)
			return
		}
		ctx.Status(204)
		ctx.Send()
	})
}
