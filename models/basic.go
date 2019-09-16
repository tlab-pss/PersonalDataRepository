package models

import "time"

type Basic struct {
	ID string
	Name string
	Birthday time.Time
	Gender int
	Mail string
	Weight string // TODO: 将来的に身体関係のテーブルができたらそっちに引っ越し
	Height string // TODO: 将来的に身体関係のテーブルができたらそっちに引っ越し
	CreatedAt time.Time
}
