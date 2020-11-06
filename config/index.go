package config

import (
	"backend/domains"
	"backend/routes"
	"backend/services"
	"github.com/gofiber/fiber"
)

func Configure(app *fiber.App) error {
	db, err := GetConnection()
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&domains.Colaborador{}, &domains.Vaga{}, &domains.VagaColaborador{})
	if err != nil {
		return err
	}

	colaboradorService := services.ColaboradorServiceFactory(db)
	vagaService := services.VagaServiceFactory(db)
	vagaColaboradorService := services.VagaColaboradorServiceFactory(db, colaboradorService, vagaService)

	routes.ConfigColaboradorRoute(colaboradorService, app)
	routes.ConfigVagaRoute(vagaService, app)
	routes.ConfigVagaColaboradorRoute(vagaColaboradorService, app)

	return nil
}
