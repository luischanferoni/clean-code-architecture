package service

import (
	"opensea/internal/ports"
)

type OpenseaService struct {
	repository ports.OpenseaRepositoryContract // interface
}

func NewService(r ports.OpenseaRepositoryContract) *OpenseaService {
	return &OpenseaService{repository: r}
}

var _ ports.OpenseaServiceContract = (*OpenseaService)(nil)
