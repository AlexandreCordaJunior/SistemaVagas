package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Vaga struct {
	gorm.Model
	Nome            string            `json:"nome"`
	Quantidade      string            `json:"quantidade"`
	Tipo            string            `json:"tipo"` // interna ou externa
	Salario         string            `json:"salario"`
	Descricao       string            `json:"descricao"`
	VagaColaborador []VagaColaborador `json:"Vaga" gorm:"foreignKey:VagaId"`
}

func (v *Vaga) Validate() error {
	if len(v.Nome) == 0 {
		return errors.New("vaga inválido")
	}
	return nil
}

//DTO
type VagaSimplesDTO struct {
	Nome       string `json:"nome"`
	Quantidade string `json:"quantidade"`
	Tipo       string `json:"tipo"` // interna ou externa
	Salario    string `json:"salario"`
	Descricao  string `json:"descricao"`
}

func (v *Vaga) GetVagaSimplesDTO() *VagaSimplesDTO {
	return &VagaSimplesDTO{
		Nome:       v.Nome,
		Quantidade: v.Quantidade,
		Tipo:       v.Tipo,
		Salario:    v.Salario,
		Descricao:  v.Descricao,
	}
}

func (dto *VagaSimplesDTO) FromDTOSimples() *Vaga {
	return &Vaga{
		Nome:       dto.Nome,
		Quantidade: dto.Quantidade,
		Tipo:       dto.Tipo,
		Salario:    dto.Salario,
		Descricao:  dto.Descricao,
	}
}

type VagaComplexaDTO struct {
	Nome            string                          `json:"nome"`
	Quantidade      string                          `json:"quantidade"`
	Tipo            string                          `json:"tipo"` // interna ou externa
	Salario         string                          `json:"salario"`
	Descricao       string                          `json:"descricao"`
	VagaColaborador []VagaColaboradorSimpleParaVaga `json:"vagaColaborador"`
}

func (v *Vaga) GetVagaComplexaDTO() *VagaComplexaDTO {
	var vagaColaboradorDTO []VagaColaboradorSimpleParaVaga
	for _, c := range v.VagaColaborador {
		vagaColaboradorDTO = append(vagaColaboradorDTO, c.getVagaColaboradorSimpleParaVaga())
	}

	return &VagaComplexaDTO{
		Nome:            v.Nome,
		Quantidade:      v.Quantidade,
		Tipo:            v.Tipo,
		Salario:         v.Salario,
		Descricao:       v.Descricao,
		VagaColaborador: vagaColaboradorDTO,
	}
}
