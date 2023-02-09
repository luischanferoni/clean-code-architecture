package postgresrepository

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"opensea/internal/domain"
)

func (r *OpenseaPostgresRepository) GetAllMovie(
	ctx context.Context,
	page int,
) ([]*domain.Movie, error) {
	var movies []*domain.Movie
	offset := (page - 1) * 10

	err := r.Database.NewSelect().
		Model(&movies).
		Relation("Creator").
		Relation("CoCreators").
		OrderExpr("created_at DESC").
		Offset(offset).
		Limit(10).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (r *OpenseaPostgresRepository) GetMovieById(ctx context.Context, movieId int64) (*domain.Movie, error) {

	movie := new(domain.Movie)
	err := r.Database.NewSelect().
		Model(movie).
		Relation("Creator").
		Relation("CoCreators").
		Where("n.id =?", movieId).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *OpenseaPostgresRepository) InsertMovie(ctx *gin.Context, movie *domain.Movie) error {

	tx, err := r.Database.BeginTx(ctx, nil)
	fail := func(err error) error {
		tx.Rollback()
		return fmt.Errorf("Buy: %v", err)
	}

	if err != nil {
		return fail(err)
	}

	_, err = tx.NewInsert().Model(movie).Exec(ctx)
	if err != nil {
		return fail(err)
	}

	log.Printf("%+v: ", movie)
	if len(movie.CoCreators) == 0 {
		// Commit the transaction.
		if err = tx.Commit(); err != nil {
			return fail(err)
		}
		return nil
	}

	var movie_cocreators []domain.MovieToCocreator
	for _, cocreator := range movie.CoCreators {
		movie_cocreators = append(movie_cocreators, domain.MovieToCocreator{
			MovieID:             movie.ID,
			StreamingPlatformID: cocreator.ID})
	}

	_, err = tx.NewInsert().Model(&movie_cocreators).Exec(ctx)
	if err != nil {
		return fail(err)
	}
	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return nil
}

func (r *OpenseaPostgresRepository) UpdateMovie(ctx context.Context, movie *domain.Movie) (*domain.Movie, error) {
	_, err := r.Database.NewUpdate().Model(movie).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (r *OpenseaPostgresRepository) UpdateStreamingPlatform(ctx context.Context, user *domain.StreamingPlatform) (*domain.StreamingPlatform, error) {
	_, err := r.Database.NewUpdate().Model(user).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *OpenseaPostgresRepository) Buy(
	ctx context.Context,
	buyer *domain.StreamingPlatform,
	creator *domain.StreamingPlatform,
	coCreators []*domain.StreamingPlatform,
) error {

	tx, err := r.Database.BeginTx(ctx, nil)
	fail := func(err error) error {
		defer tx.Rollback()
		return fmt.Errorf("Buy: %v", err)
	}

	if err != nil {
		return fail(err)
	}

	// 2. Update buyer balance
	_, err = tx.NewUpdate().Model(buyer).Column("balance").WherePK().Exec(ctx)
	if err != nil {
		return fail(err)
	}

	// 3. Update creator balance
	_, err = tx.NewUpdate().Model(creator).Column("balance").WherePK().Exec(ctx)
	if err != nil {
		return fail(err)
	}

	// b) Update Co-creators balance
	_, err = tx.NewUpdate().
		Model(&coCreators).
		Column("balance").
		Bulk().
		WherePK().
		Exec(ctx)
	if err != nil {
		return fail(err)
	}

	// Commit the transaction.
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return nil
}
