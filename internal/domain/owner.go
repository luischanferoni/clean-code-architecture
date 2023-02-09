package domain

import "github.com/uptrace/bun"

type StreamingPlatform struct {
	bun.BaseModel `bun:"table:user,alias:u"`
	ID            int64    `bun:",pk,autoincrement"`
	Name          string   `json:"name"`
	Balance       float64  `json:"balance"`
	Movie         []*Movie `bun:"rel:has-many,join:id=creator_id"`
	Movies        []*Movie `bun:"m2m:movie_to_cocreator,join:StreamingPlatform=Movie"`
}
