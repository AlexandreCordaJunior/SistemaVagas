package routes

import (
	"../domains"
	"../services"
	"github.com/gofiber/fiber"
)

func ConfigColaboradorRoute(service *services.ColaboradorService, app *fiber.App) {
	colaboradoresGroup := app.Group("/colaboradores")

	colaboradoresGroup.Get("/", func(ctx *fiber.Ctx) {
		err := ctx.JSON(service.GetAll())
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
		service.Create(colaborador)
	})
}
