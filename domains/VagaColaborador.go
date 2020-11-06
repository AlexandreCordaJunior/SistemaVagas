package domains

import (
	"gorm.io/gorm"
)

type VagaColaborador struct {
	gorm.Model
	Colaborador   Colaborador `gorm:"foreignKey:ColaboradorId"`
	ColaboradorId int64
	Vaga          Vaga `gorm:"foreignKey:VagaId"`
	VagaId        int64
}

type VagaColaboradorSimpleParaVaga struct {
	Colaborador ColaboradorSimplesDTO
}

func (vagaColaborador *VagaColaborador) getVagaColaboradorSimpleParaVaga() VagaColaboradorSimpleParaVaga {
	return VagaColaboradorSimpleParaVaga{
		Colaborador: *vagaColaborador.Colaborador.GetColaboradorSimplesDTO(),
	}
}

type VagaColaboradorSimpleParaColaborador struct {
	Vaga VagaSimplesDTO
}

func (vagaColaborador *VagaColaborador) getVagaColaboradorSimpleParaColaborador() VagaColaboradorSimpleParaColaborador {
	return VagaColaboradorSimpleParaColaborador{
		Vaga: *vagaColaborador.Vaga.GetVagaSimplesDTO(),
	}
}

type VagaColaboradorDTO struct {
	Vaga        VagaSimplesDTO        `json:"vaga"`
	Colaborador ColaboradorSimplesDTO `json:"colaborador"`
}

func (vagaColaborador *VagaColaborador) GetVagaColaboradorDTO() *VagaColaboradorDTO {
	return &VagaColaboradorDTO{
		Vaga:        *vagaColaborador.Vaga.GetVagaSimplesDTO(),
		Colaborador: *vagaColaborador.Colaborador.GetColaboradorSimplesDTO(),
	}
}

type VagaColaboradorEntrada struct {
	gorm.Model
	VagaId        int64
	ColaboradorId int64
}
