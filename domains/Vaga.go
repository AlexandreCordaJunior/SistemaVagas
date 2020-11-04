package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Vaga struct {
	gorm.Model
	Nome       string `json:"nome"`
	Quantidade string `json:"quantidade"`
	Tipo       string `json:"tipo"` // interna ou externa
	Salario    string `json:"salario"`
	Descricao  string `json:"descricao"`
}

func (v *Vaga) Validate() error {
	if len(v.Nome) == 0 {
		return errors.New("vaga inválido")
	}
	return nil
}
