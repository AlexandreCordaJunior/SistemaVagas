package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Colaborador struct {
	gorm.Model
	Nome  string  `json:"nome"`
	Idade int     `json:"idade"`
	Email string  `json:"email"`
	Hash  string  `json:"hash"`
	Vaga  []*Vaga `json:"vaga" gorm:"many2many:vaga_colaborador;"`
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
	Nome  string            `json:"nome"`
	Idade int               `json:"idade"`
	Email string            `json:"email"`
	Hash  string            `json:"hash"`
	Vaga  []*VagaSimplesDTO `json:"vaga"`
}

func (c *Colaborador) GetColaboradorComplexoDTO() *ColaboradorComplexoDTO {
	var vagaDTO []*VagaSimplesDTO
	for _, v := range c.Vaga {
		vagaDTO = append(vagaDTO, v.GetVagaSimplesDTO())
	}

	return &ColaboradorComplexoDTO{
		Nome:  c.Nome,
		Idade: c.Idade,
		Email: c.Email,
		Vaga:  vagaDTO,
	}
}
