package domain

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Movie struct {
	bun.BaseModel `bun:"table:movie,alias:n"`
	ID            int64                `bun:",pk,autoincrement" json:"id"`
	File          string               `form:"file" binding:"required"`
	Description   string               `form:"description" binding:"required"`
	CreatorID     int64                `form:"creator_id" binding:"required"`
	Creator       *StreamingPlatform   `bun:"rel:belongs-to,join:creator_id=id"`
	Price         float64              `form:"price" binding:"required"`
	CoCreators    []*StreamingPlatform `bun:"m2m:movie_to_cocreator,join:Movie=StreamingPlatform"`
	// metadata
	CreatedAt int64 `json:"created_at"`
	CreatedBy int64 `json:"created_by"`
	UpdatedAt int64 `json:"updated_at"`
	UpdatedBy int64 `json:"updated_by"`
}

var _ bun.BeforeAppendModelHook = (*Movie)(nil)

func (m *Movie) BeforeAppendModel(ctx context.Context, query bun.Query) error {
	switch query.(type) {
	case *bun.InsertQuery:
		m.CreatedAt = time.Now().Unix()
	case *bun.UpdateQuery:
		m.UpdatedAt = time.Now().Unix()
	}
	return nil
}

type MovieToCocreator struct {
	bun.BaseModel       `bun:"table:movie_to_cocreator,alias:n_t_c"`
	MovieID             int64              `bun:",pk"`
	Movie               *Movie             `bun:"rel:belongs-to,join:movie_id=id"`
	StreamingPlatformID int64              `bun:",pk"`
	StreamingPlatform   *StreamingPlatform `bun:"rel:belongs-to,join:user_id=id"`
}
