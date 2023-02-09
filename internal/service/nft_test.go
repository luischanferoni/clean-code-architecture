package service

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"

	"opensea/internal/domain"
	mock_ports "opensea/mocks"
)

type mocks struct {
	repository *mock_ports.MockOpenseaRepositoryContract
}

func newMocks() *mocks {
	return &mocks{
		repository: new(mock_ports.MockOpenseaRepositoryContract),
	}
}

func newSubject(mocks *mocks) *OpenseaService {
	return &OpenseaService{repository: mocks.repository}
}

func TestCreateShouldSuccess(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"valid request": testValid,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testValid(t *testing.T) {
	mocks := newMocks()
	subject := newSubject(mocks)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	movie := &domain.Movie{
		File:        "my_image.JPEG",
		Description: "My token",
		CreatorID:   1,
		Price:       10.00,
	}
	/*
		user := domain.StreamingPlatform{
			ID:      1,
			Name:    "movie creator 1",
			Balance: 100.00,
		}
	*/
	mocks.repository.EXPECT().InsertMovie(ctx, movie).Return(nil)
	err := subject.Create(ctx, movie)
	require.Nil(t, err)
}

/*
func TestCreateShouldFail(t *testing.T) {
	for scenario, fn := range map[string]func(t *testing.T){
		"incomplete values for movie": testMissingDescription,
	} {
		t.Run(scenario, func(t *testing.T) {
			fn(t)
		})
	}
}

func testMissingDescription(t *testing.T) {
	mocks := newMocks()
	subject := newSubject(mocks)
	ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
	movie := &domain.Movie{
		File:      "my_image.JPEG",
		CreatorID: 1,
		Price:     10.00,
	}
	mocks.repository.EXPECT().InsertMovie(ctx, movie).Return(nil, errors.New("Error"))
	_, err := subject.Create(ctx, movie)
	require.Error(t, err)
}
*/
