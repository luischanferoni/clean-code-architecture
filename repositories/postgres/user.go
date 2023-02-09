package postgresrepository

import (
	"context"
	//"errors"

	"opensea/internal/domain"
)

func (r *OpenseaPostgresRepository) GetStreamingPlatformById(ctx context.Context, userId int64) (*domain.StreamingPlatform, error) {

	user := new(domain.StreamingPlatform)
	err := r.Database.NewSelect().Model(user).Where("u.id =?", userId).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}
