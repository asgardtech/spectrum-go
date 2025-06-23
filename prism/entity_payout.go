package prism

import "time"

type Payout struct {
	ID          string
	Amount      float64
	Currency    string
	Description string
	Recipient   string
	Status      string
	CreatedAt   time.Time
	ProcessedAt *time.Time
	Selected    bool
}
