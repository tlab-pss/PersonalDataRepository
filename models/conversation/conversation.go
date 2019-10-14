package conversation

import "time"

type Conversation struct {
	ID            string
	TransactionID string
	RequestText   string
	ResponseText  string
	CreatedAt     time.Time
}
