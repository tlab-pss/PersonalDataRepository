package hotpepper

import "time"

// todo: とりあえず以下。なんかあったら足す
type Intake struct {
	ID              string
	Menu            string   // Vision AIのWeb Entitiesで最もポイントが高かったやつ
	Photo           string   // base64
	Calorie         float64  // unit: kcal
	Labels          []string // Vision AIのLabelsで80%以上のやつ
	SmallCategoryID string
	CreatedAt       time.Time
}
