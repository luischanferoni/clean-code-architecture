package service

import (
	"errors"
	"log"

	"github.com/gin-gonic/gin"

	"opensea/internal/domain"
	"opensea/shared"
)

func (s *OpenseaService) Create(ctx *gin.Context, movie *domain.Movie) error {

	err := s.repository.InsertMovie(ctx, movie)
	if err != nil {
		return err
	}

	return nil
}

func (s *OpenseaService) Get(ctx *gin.Context, movieId int64) (*domain.Movie, error) {

	movie, err := s.repository.GetMovieById(ctx, movieId)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (s *OpenseaService) GetAll(ctx *gin.Context, page int) ([]*domain.Movie, error) {

	movies, err := s.repository.GetAllMovie(ctx, page)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (s *OpenseaService) Buy(ctx *gin.Context, movieId int64) error {
	// get the buyer id from context
	buyerId := shared.GetContextStreamingPlatformId(ctx)
	// 1. Get the buyer
	buyer, err := s.repository.GetStreamingPlatformById(ctx, buyerId)
	if err != nil {
		return err
	}

	// 2. Get the Movie
	movie, err := s.repository.GetMovieById(ctx, movieId)
	if err != nil {
		return err
	}

	// 3. check if the user has balance to buy this movie
	if movie.Price > buyer.Balance {
		return errors.New("What's up?")
	}

	// 4. Buy
	// a) Substract from the buyer balance
	buyer.Balance = buyer.Balance - movie.Price

	// b) Fee the creator
	creator := movie.Creator
	creator.Balance = creator.Balance + movie.Price*0.8 // 80% to creator

	// b) Fee to co-creators
	restOfBalance := movie.Price * 0.2                                    // 20% to the co-creators
	percentageCoCreator := restOfBalance / float64(len(movie.CoCreators)) // divided equaly

	var coCreators []*domain.StreamingPlatform
	for _, cocreator := range movie.CoCreators {
		cocreator.Balance = cocreator.Balance + percentageCoCreator
		coCreators = append(coCreators, cocreator)
	}
	log.Printf("%+v", buyer)
	log.Printf("%+v", creator)
	log.Printf("%+v", &coCreators)
	err = s.repository.Buy(ctx, buyer, creator, coCreators)

	return nil

}
