package services

import (
	"backend/domains"
	"gorm.io/gorm"
)

type ColaboradorService struct {
	db *gorm.DB
}

func ColaboradorServiceFactory(db *gorm.DB) *ColaboradorService {
	return &ColaboradorService{db}
}

func (s *ColaboradorService) GetAll() ([]*domains.Colaborador, error) {
	var colaboradorVet []*domains.Colaborador
	err := s.db.Find(&colaboradorVet).Error
	if err != nil {
		return nil, err
	}
	return colaboradorVet, nil
}

func (s *ColaboradorService) GetOne(id string) (*domains.Colaborador, error) {
	var colaborador *domains.Colaborador
	err := s.db.Where("id = ?", id).First(colaborador).Error
	if err != nil {
		return nil, err
	}
	return colaborador, nil
}

func (s *ColaboradorService) Create(colaborador *domains.Colaborador) (*domains.Colaborador, error) {
	err := colaborador.Validate()
	if err != nil {
		return nil, err
	}
	s.db.Create(colaborador)
	return colaborador, nil
}
