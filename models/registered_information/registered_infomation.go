package registered_information

import "time"

// TODO: 認証情報はここに入れる

type RegisteredInformation struct {
	ID        string
	Mail      string
	CreatedAt time.Time
}
