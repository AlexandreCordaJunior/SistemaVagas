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
	colaboradorService := services.ColaboradorServiceFactory(db)
	routes.ConfigColaboradorRoute(colaboradorService, app)

	return nil
}
