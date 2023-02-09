package restcontroller

import (
	"opensea/internal/ports"
)

type RestController struct {
	service ports.OpenseaServiceContract // interface
}

func NewController(s ports.OpenseaServiceContract) *RestController {
	return &RestController{service: s}
}
