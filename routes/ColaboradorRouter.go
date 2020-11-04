package routes

import (
	"backend/domains"
	"backend/services"
	"github.com/gofiber/fiber"
	"strconv"
)

func ConfigColaboradorRoute(service *services.ColaboradorService, app *fiber.App) {
	colaboradoresGroup := app.Group("/colaboradores")

	colaboradoresGroup.Get("/", func(ctx *fiber.Ctx) {
		colaborador, err := service.GetAll()
		if err != nil {
			ctx.Next(err)
			return
		}
		err = ctx.JSON(colaborador)
		if err != nil {
			ctx.Next(err)
		}
	})

	colaboradoresGroup.Get("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		colaborador, err := service.GetOne(id)
		if err != nil {
			ctx.Next(err)
			return
		}
		err = ctx.JSON(colaborador)
		if err != nil {
			ctx.Next(err)
		}
	})

	colaboradoresGroup.Post("/", func(ctx *fiber.Ctx) {
		colaborador := &domains.Colaborador{}
		if err := ctx.BodyParser(colaborador); err != nil {
			ctx.Next(err)
			return
		}
		colaborador, err := service.Create(colaborador)
		if err != nil {
			ctx.Next(err)
		}
		ctx.Set("Location", ctx.Path()+strconv.Itoa(int(colaborador.ID)))
		ctx.Status(201)
		ctx.Send()
	})

	colaboradoresGroup.Put("/:id", func(ctx *fiber.Ctx) {
		id := ctx.Params("id")
		colaborador := &domains.Colaborador{}
		if err := ctx.BodyParser(colaborador); err != nil {
			ctx.Next(err)
			return
		}

		colaborador, err := service.Update(colaborador, id)
		if err != nil {
			ctx.Next(err)
			return
		}

		err = ctx.JSON(colaborador)
		if err != nil {
			ctx.Next(err)
		}
	})

	colaboradoresGroup.Delete("/:id", func(ctx *fiber.Ctx) {
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
