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

	//Colaboradores
	err = db.AutoMigrate(&domains.Colaborador{})
	if err != nil {
		return err
	}

	//Vagas
	err = db.AutoMigrate(&domains.Vaga{})
	if err != nil {
		return err
	}

	colaboradorService := services.ColaboradorServiceFactory(db)
	vagaService := services.VagaServiceFactory(db)

	routes.ConfigColaboradorRoute(colaboradorService, app)
	routes.ConfigVagaRoute(vagaService, app)

	return nil
}
