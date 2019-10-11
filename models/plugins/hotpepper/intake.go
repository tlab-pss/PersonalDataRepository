package hotpepper

import "time"

type Intake struct {
	ID        string
	Menu      string
	Calorie    float64
	CreatedAt time.Time
}
