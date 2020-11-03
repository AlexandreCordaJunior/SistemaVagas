package domains

import (
	"errors"
	"gorm.io/gorm"
)

type Colaborador struct {
	gorm.Model
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

func (c *Colaborador) Validate() error {
	if len(c.Nome) == 0 {
		return errors.New("colaborador inválido")
	}
	return nil
}
