package entities

import "time"

type Transaction struct {
	// RecordID          int
	CustomerRef       string
	AgentCode         string
	AgentID           string
	Amount            float32
	PaymentDate       time.Time
	CompletedAtVendor bool
	SentToUtility     bool
	UtilitySentDate   time.Time
	UtilityReference  string
	RecordDate        time.Time
	// Affected   map[string]string
}
