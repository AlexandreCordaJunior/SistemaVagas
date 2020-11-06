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
	colaborador := &domains.Colaborador{}
	colaborador.VagaColaborador = make([]domains.VagaColaborador, 0)
	rows, err := s.db.Table("colaboradors").Where("colaboradors.id = ?", id).
		Joins("Join vaga_colaboradors on vaga_colaboradors.colaborador_id = colaboradors.id").
		Joins("Join vagas on vaga_colaboradors.vaga_id = vagas.id").
		Select("colaboradors.id, colaboradors.nome, colaboradors.idade, colaboradors.hash, colaboradors.email, " +
			"vaga_colaboradors.id, vaga_colaboradors.vaga_id, vaga_colaboradors.colaborador_id, " +
			"vagas.id, vagas.nome, vagas.quantidade, vagas.tipo, vagas.salario, vagas.descricao").Rows()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		vagaColaborador := domains.VagaColaborador{}
		vaga := domains.Vaga{}
		err = rows.Scan(&colaborador.ID, &colaborador.Nome, &colaborador.Idade, &colaborador.Hash, &colaborador.Email, &vagaColaborador.ID,
			&vagaColaborador.VagaId, &vagaColaborador.ColaboradorId, &vaga.ID, &vaga.Nome, &vaga.Quantidade, &vaga.Tipo,
			&vaga.Salario, &vaga.Descricao)
		if err != nil {
			return nil, err
		}
		vagaColaborador.Colaborador = colaborador
		vagaColaborador.Vaga = &vaga
		colaborador.VagaColaborador = append(colaborador.VagaColaborador, vagaColaborador)
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

func (s *ColaboradorService) Update(colaboradorNovo *domains.Colaborador, id string) (*domains.Colaborador, error) {
	colaboradorAntigo, err := s.GetOne(id)
	if err != nil {
		return nil, err
	}

	UpdateColaborador(colaboradorAntigo, colaboradorNovo)
	err = s.db.Save(colaboradorNovo).Error
	if err != nil {
		return nil, err
	}

	return colaboradorNovo, nil
}

func (s *ColaboradorService) Delete(id string) error {
	colaborador := &domains.Colaborador{}
	err := s.db.Delete(colaborador, id).Error
	if err != nil {
		return err
	}
	return nil
}
