package small_category

import (
	"time"
)

type SmallCategory struct {
	ID            string
	Name          string
	BigCategoryId string
	CreatedAt     time.Time
}
