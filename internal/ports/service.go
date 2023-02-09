package ports

import (
	"github.com/gin-gonic/gin"

	"opensea/internal/domain"
)

//go:generate mockgen -source=service.go -destination=../../mocks/mock_service.go
type OpenseaServiceContract interface {
	Create(ctx *gin.Context, movie *domain.Movie) error
	Get(ctx *gin.Context, movieId int64) (*domain.Movie, error)
	GetAll(ctx *gin.Context, page int) ([]*domain.Movie, error)
	Buy(ctx *gin.Context, movieId int64) error
}
