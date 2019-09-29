package basic

import "time"

type Basic struct {
	ID        string
	Name      string
	Birthday  time.Time
	Gender    int
	CreatedAt time.Time
}
