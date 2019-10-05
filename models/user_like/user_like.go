package user_like

import (
	"time"
)

type UserLike struct {
	ID              string
	SmallCategoryId string
	CreatedAt       time.Time
}
