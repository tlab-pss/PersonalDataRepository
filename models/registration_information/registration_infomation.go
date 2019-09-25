package registration_information

import "time"

// TODO: 認証情報はここに入れる

type RegistrationInformation struct {
	ID        string
	Mail      string
	CreatedAt time.Time
}
