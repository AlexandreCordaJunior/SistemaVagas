package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Colaborador struct {
	gorm.Model
	Nome            string            `json:"nome"`
	Idade           int               `json:"idade"`
	Email           string            `json:"email"`
	Hash            string            `json:"hash"`
	VagaColaborador []VagaColaborador `json:"Vaga" gorm:"foreignKey:ColaboradorId"`
}

func (c *Colaborador) Validate() error {
	if len(c.Nome) == 0 {
		return errors.New("colaborador inválido")
	}
	return nil
}

//DTO
type ColaboradorSimplesDTO struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Email string `json:"email"`
}

func (c *Colaborador) GetColaboradorSimplesDTO() *ColaboradorSimplesDTO {
	return &ColaboradorSimplesDTO{
		Nome:  c.Nome,
		Idade: c.Idade,
		Email: c.Email,
	}
}

func (dto *ColaboradorSimplesDTO) FromDTOSimples() *Colaborador {
	return &Colaborador{
		Nome:  dto.Nome,
		Idade: dto.Idade,
		Email: dto.Email,
	}
}

type ColaboradorComplexoDTO struct {
	Nome            string                                 `json:"nome"`
	Idade           int                                    `json:"idade"`
	Email           string                                 `json:"email"`
	VagaColaborador []VagaColaboradorSimpleParaColaborador `json:"vagaColaborador"`
}

func (c *Colaborador) GetColaboradorComplexoDTO() *ColaboradorComplexoDTO {
	var vagaDto []VagaColaboradorSimpleParaColaborador

	for _, v := range c.VagaColaborador {
		vagaDto = append(vagaDto, v.getVagaColaboradorSimpleParaColaborador())
	}

	return &ColaboradorComplexoDTO{
		Nome:            c.Nome,
		Idade:           c.Idade,
		Email:           c.Email,
		VagaColaborador: vagaDto,
	}
}
