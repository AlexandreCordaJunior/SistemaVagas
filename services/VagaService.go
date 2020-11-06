package services

import (
	"backend/domains"
	"gorm.io/gorm"
)

type VagaService struct {
	db *gorm.DB
}

func VagaServiceFactory(db *gorm.DB) *VagaService {
	return &VagaService{db}
}

func (s *VagaService) GetAll() ([]*domains.Vaga, error) {
	var vagaVet []*domains.Vaga
	err := s.db.Find(&vagaVet).Error
	if err != nil {
		return nil, err
	}
	return vagaVet, nil
}

func (s *VagaService) GetOne(id string) (*domains.Vaga, error) {
	vaga := &domains.Vaga{}
	vaga.VagaColaborador = make([]domains.VagaColaborador, 0)
	rows, err := s.db.Table("vagas").Where("vagas.id = ?", id).
		Joins("Join vaga_colaboradors on vaga_colaboradors.vaga_id = vagas.id").
		Joins("Join colaboradors on vaga_colaboradors.colaborador_id = colaboradors.id").
		Select("vagas.id, vagas.nome, vagas.quantidade, vagas.tipo, vagas.salario, vagas.descricao, " +
			"colaboradors.id, colaboradors.nome, colaboradors.idade, colaboradors.hash, colaboradors.email, " +
			"vaga_colaboradors.id, vaga_colaboradors.vaga_id, vaga_colaboradors.colaborador_id").Rows()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		vagaColaborador := domains.VagaColaborador{}
		colaborador := domains.Colaborador{}

		err = rows.Scan(&vaga.ID, &vaga.Nome, &vaga.Quantidade, &vaga.Tipo, &vaga.Salario, &vaga.Descricao,
			&colaborador.ID, &colaborador.Nome, &colaborador.Idade, &colaborador.Hash, &colaborador.Email,
			&vagaColaborador.ID, &vagaColaborador.VagaId, &vagaColaborador.ColaboradorId)
		if err != nil {
			return nil, err
		}
		vagaColaborador.Colaborador = &colaborador
		vagaColaborador.Vaga = vaga
		vaga.VagaColaborador = append(vaga.VagaColaborador, vagaColaborador)
	}
	return vaga, nil
}

func (s *VagaService) Create(vaga *domains.Vaga) (*domains.Vaga, error) {
	err := vaga.Validate()
	if err != nil {
		return nil, err
	}
	s.db.Create(vaga)
	return vaga, nil
}

func (s *VagaService) Update(vagaNovo *domains.Vaga, id string) (*domains.Vaga, error) {
	vagaAntigo, err := s.GetOne(id)
	if err != nil {
		return nil, err
	}

	UpdateVaga(vagaAntigo, vagaNovo)
	err = s.db.Save(vagaNovo).Error
	if err != nil {
		return nil, err
	}

	return vagaNovo, nil
}

func (s *VagaService) Delete(id string) error {
	vaga := &domains.Vaga{}
	err := s.db.Delete(vaga, id).Error
	if err != nil {
		return err
	}
	return nil
}
