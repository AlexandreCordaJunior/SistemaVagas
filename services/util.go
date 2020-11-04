package services

import "backend/domains"

func UpdateColaborador(colaboradorAntigo *domains.Colaborador, colaboradorNovo *domains.Colaborador) {
	colaboradorNovo.ID = colaboradorAntigo.ID
	colaboradorNovo.CreatedAt = colaboradorAntigo.CreatedAt
	colaboradorNovo.DeletedAt = colaboradorAntigo.DeletedAt
	colaboradorNovo.UpdatedAt = colaboradorAntigo.UpdatedAt
	colaboradorNovo.Hash = colaboradorAntigo.Hash
}

func UpdateVaga(colaboradorAntigo *domains.Vaga, colaboradorNovo *domains.Vaga) {
	colaboradorNovo.ID = colaboradorAntigo.ID
	colaboradorNovo.CreatedAt = colaboradorAntigo.CreatedAt
	colaboradorNovo.DeletedAt = colaboradorAntigo.DeletedAt
	colaboradorNovo.UpdatedAt = colaboradorAntigo.UpdatedAt
}
