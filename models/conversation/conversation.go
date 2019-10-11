package conversation

import "time"

type Conversation struct {
	ID            string
	TransactionId string
	RequestText   string
	ResponseText  string
	CreatedAt     time.Time
}
