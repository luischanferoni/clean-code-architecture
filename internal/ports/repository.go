package ports

import (
	"context"

	"github.com/gin-gonic/gin"

	"opensea/internal/domain"
)

//go:generate mockgen -source=repository.go -destination=../../mocks/mock_repository.go
type OpenseaRepositoryContract interface {
	GetAllMovie(ctx context.Context, page int) ([]*domain.Movie, error)
	GetMovieById(ctx context.Context, movieId int64) (*domain.Movie, error)
	InsertMovie(ctx *gin.Context, movie *domain.Movie) error
	UpdateMovie(ctx context.Context, movie *domain.Movie) (*domain.Movie, error)
	GetStreamingPlatformById(ctx context.Context, userId int64) (*domain.StreamingPlatform, error)
	UpdateStreamingPlatform(ctx context.Context, user *domain.StreamingPlatform) (*domain.StreamingPlatform, error)
	Buy(ctx context.Context, buyer *domain.StreamingPlatform, creator *domain.StreamingPlatform, coCreators []*domain.StreamingPlatform) error
}
