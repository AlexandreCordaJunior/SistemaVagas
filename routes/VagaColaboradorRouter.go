package routes

import (
	"backend/domains"
	"backend/services"
	"github.com/gofiber/fiber"
	"strconv"
)

func ConfigVagaColaboradorRoute(service *services.VagaColaboradorService, app *fiber.App) {
	vagaColaboradorsGroup := app.Group("/vagaColaboradores")

	vagaColaboradorsGroup.Get("/", func(ctx *fiber.Ctx) {
		vagaColaborador, err := service.GetAll()
		if err != nil {
			ctx.Next(err)
			return
		}
		var vagaColaboradorDTO []*domains.VagaColaboradorDTO
		for _, v := range vagaColaborador {
			vagaColaboradorDTO = append(vagaColaboradorDTO, v.GetVagaColaboradorDTO())
		}
		err = ctx.JSON(vagaColaboradorDTO)
		if err != nil {
			ctx.Next(err)
		}
	})

	vagaColaboradorsGroup.Get("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		vagaColaborador, err := service.GetOne(id)
		if err != nil {
			ctx.Next(err)
			return
		}
		vagaColaboradorDTO := vagaColaborador.GetVagaColaboradorDTO()
		err = ctx.JSON(vagaColaboradorDTO)
		if err != nil {
			ctx.Next(err)
		}
	})

	vagaColaboradorsGroup.Post("/", func(ctx *fiber.Ctx) {
		vagaColaboradorDTO := &domains.VagaColaboradorEntrada{}
		if err := ctx.BodyParser(vagaColaboradorDTO); err != nil {
			ctx.Next(err)
			return
		}

		vagaColaborador, err := service.Create(vagaColaboradorDTO)
		if err != nil {
			ctx.Next(err)
		}
		ctx.Set("Location", ctx.Path()+strconv.Itoa(int(vagaColaborador.ID)))
		ctx.Status(201)
		ctx.Send()
	})

	//vagaColaboradorsGroup.Put("/:id", func(ctx *fiber.Ctx) {
	//	id := ctx.Params("id")
	//	vagaColaboradorDTO := &domains.VagaColaboradorSimplesDTO{}
	//	if err := ctx.BodyParser(vagaColaboradorDTO); err != nil {
	//		ctx.Next(err)
	//		return
	//	}
	//
	//	vagaColaborador := vagaColaboradorDTO.FromDTOSimples()
	//	vagaColaborador, err := service.Update(vagaColaborador, id)
	//	if err != nil {
	//		ctx.Next(err)
	//		return
	//	}
	//
	//	err = ctx.JSON(vagaColaborador)
	//	if err != nil {
	//		ctx.Next(err)
	//	}
	//})

	vagaColaboradorsGroup.Delete("/:id", func(ctx *fiber.Ctx) {
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
