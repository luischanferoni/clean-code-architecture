package postgresrepository

import (
	"context"
	"opensea/internal/domain"
	"opensea/internal/ports"

	"github.com/uptrace/bun"
)

// struct and its corresponing new method
type OpenseaPostgresRepository struct {
	Database *bun.DB
}

func NewPostgresRepository(db *bun.DB) *OpenseaPostgresRepository {
	return &OpenseaPostgresRepository{Database: db}
}

// checking if the struct complies with the contract (interface)
var _ ports.OpenseaRepositoryContract = (*OpenseaPostgresRepository)(nil)

// Just a necesary step for the many to many relationships
func RegisterModel(ctx context.Context, db *bun.DB) {

	db.RegisterModel((*domain.MovieToCocreator)(nil))

}
