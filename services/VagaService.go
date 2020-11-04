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
	err := s.db.First(vaga, id).Error
	if err != nil {
		return nil, err
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
