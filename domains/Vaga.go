package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Vaga struct {
	gorm.Model
	Nome        string         `json:"nome"`
	Quantidade  string         `json:"quantidade"`
	Tipo        string         `json:"tipo"` // interna ou externa
	Salario     string         `json:"salario"`
	Descricao   string         `json:"descricao"`
	Colaborador []*Colaborador `json:"colaborador" gorm:"many2many:vaga_colaborador;"`
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
	Nome        string                   `json:"nome"`
	Quantidade  string                   `json:"quantidade"`
	Tipo        string                   `json:"tipo"` // interna ou externa
	Salario     string                   `json:"salario"`
	Descricao   string                   `json:"descricao"`
	Colaborador []*ColaboradorSimplesDTO `json:"colaborador"`
}

func (v *Vaga) GetVagaComplexaDTO() *VagaComplexaDTO {
	var colaboradorDTO []*ColaboradorSimplesDTO
	for _, c := range v.Colaborador {
		colaboradorDTO = append(colaboradorDTO, c.GetColaboradorSimplesDTO())
	}

	return &VagaComplexaDTO{
		Nome:        v.Nome,
		Quantidade:  v.Quantidade,
		Tipo:        v.Tipo,
		Salario:     v.Salario,
		Descricao:   v.Descricao,
		Colaborador: colaboradorDTO,
	}
}
