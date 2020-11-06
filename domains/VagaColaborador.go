package domains

import (
	"gorm.io/gorm"
)

type VagaColaborador struct {
	gorm.Model
	Colaborador   *Colaborador `gorm:"foreignKey:ColaboradorId"`
	ColaboradorId int64
	Vaga          *Vaga `gorm:"foreignKey:VagaId"`
	VagaId        int64
}

type VagaColaboradorSimpleParaVaga struct {
	ID          uint                   `json:"id"`
	Colaborador *ColaboradorSimplesDTO `json:"colaborador"`
}

func (vagaColaborador *VagaColaborador) getVagaColaboradorSimpleParaVaga() VagaColaboradorSimpleParaVaga {
	return VagaColaboradorSimpleParaVaga{
		ID:          vagaColaborador.ID,
		Colaborador: vagaColaborador.Colaborador.GetColaboradorSimplesDTO(),
	}
}

type VagaColaboradorSimpleParaColaborador struct {
	ID   uint            `json:"id"`
	Vaga *VagaSimplesDTO `json:"vaga"`
}

func (vagaColaborador *VagaColaborador) getVagaColaboradorSimpleParaColaborador() VagaColaboradorSimpleParaColaborador {
	return VagaColaboradorSimpleParaColaborador{
		ID:   vagaColaborador.ID,
		Vaga: vagaColaborador.Vaga.GetVagaSimplesDTO(),
	}
}

type VagaColaboradorDTO struct {
	ID          uint                  `json:"id"`
	Vaga        VagaSimplesDTO        `json:"vaga"`
	Colaborador ColaboradorSimplesDTO `json:"colaborador"`
}

func (vagaColaborador *VagaColaborador) GetVagaColaboradorDTO() *VagaColaboradorDTO {
	return &VagaColaboradorDTO{
		ID:          vagaColaborador.ID,
		Vaga:        *vagaColaborador.Vaga.GetVagaSimplesDTO(),
		Colaborador: *vagaColaborador.Colaborador.GetColaboradorSimplesDTO(),
	}
}

type VagaColaboradorEntrada struct {
	VagaId        int64
	ColaboradorId int64
}
