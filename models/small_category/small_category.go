package small_category

import (
	"time"
)

type SmallCategory struct {
	ID            string
	Name          string
	BigCategoryID string
	CreatedAt     time.Time
}
