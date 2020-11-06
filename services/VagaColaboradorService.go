package services

import (
	"backend/domains"
	"fmt"
	"gorm.io/gorm"
)

type VagaColaboradorService struct {
	db                 *gorm.DB
	colaboradorService *ColaboradorService
	vagaService        *VagaService
}

func VagaColaboradorServiceFactory(db *gorm.DB, colaboradorService *ColaboradorService, vagaService *VagaService) *VagaColaboradorService {
	return &VagaColaboradorService{db, colaboradorService, vagaService}
}

func (s *VagaColaboradorService) GetAll() ([]*domains.VagaColaborador, error) {
	var vagaColaboradorVet []*domains.VagaColaborador
	err := s.db.Joins("Vaga").Joins("Colaborador").Find(&vagaColaboradorVet).Error
	if err != nil {
		return nil, err
	}
	return vagaColaboradorVet, nil
}

func (s *VagaColaboradorService) GetOne(id string) (*domains.VagaColaborador, error) {
	vagaColaborador := &domains.VagaColaborador{}
	err := s.db.Joins("Vaga").Joins("Colaborador").First(vagaColaborador, id).Error
	if err != nil {
		return nil, err
	}
	return vagaColaborador, nil
}

func (s *VagaColaboradorService) Create(vagaColaboradorEntrada *domains.VagaColaboradorEntrada) (*domains.VagaColaborador, error) {
	vagaColaborador := &domains.VagaColaborador{}
	vaga, err := s.vagaService.GetOne(fmt.Sprintf("%d", vagaColaboradorEntrada.VagaId))
	if err != nil {
		return nil, err
	}

	colaborador, err := s.colaboradorService.GetOne(fmt.Sprintf("%d", vagaColaboradorEntrada.ColaboradorId))
	if err != nil {
		return nil, err
	}

	vagaColaborador.Vaga = vaga
	vagaColaborador.Colaborador = colaborador

	s.db.Create(vagaColaborador)
	return vagaColaborador, nil
}

//func (s *VagaColaboradorService) Update(vagaColaboradorNovo *domains.VagaColaborador, id string) (*domains.VagaColaborador, error) {
//	vagaColaboradorAntigo, err := s.GetOne(id)
//	if err != nil {
//		return nil, err
//	}
//
//	UpdateVagaColaborador(vagaColaboradorAntigo, vagaColaboradorNovo)
//	err = s.db.Save(vagaColaboradorNovo).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return vagaColaboradorNovo, nil
//}

func (s *VagaColaboradorService) Delete(id string) error {
	vagaColaborador := &domains.VagaColaborador{}
	err := s.db.Delete(vagaColaborador, id).Error
	if err != nil {
		return err
	}
	return nil
}
