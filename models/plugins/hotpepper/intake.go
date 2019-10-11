package hotpepper

import "time"

// todo: とりあえず以下。なんかあったら足す
type Intake struct {
	ID        string
	Menu      string
	Calorie   float64
	CreatedAt time.Time
}
