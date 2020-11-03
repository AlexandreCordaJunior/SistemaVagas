package services

import (
	"../domains"
	"gorm.io/gorm"
)

type ColaboradorService struct {
	db *gorm.DB
}

func ColaboradorServiceFactory(db *gorm.DB) *ColaboradorService {
	return &ColaboradorService{db}
}

func (s *ColaboradorService) GetAll() []*domains.Colaborador {
	var colaboradorVet []*domains.Colaborador
	s.db.Find(&colaboradorVet)
	return colaboradorVet
}

func (s *ColaboradorService) Create(colaborador *domains.Colaborador) error {
	err := colaborador.Validate()
	if err != nil {
		return err
	}
	s.db.Create(colaborador)
	return nil
}
