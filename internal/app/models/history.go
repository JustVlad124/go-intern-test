package models

import "time"

type History struct {
	ID     string
	Time   time.Time
	From   string
	To     string
	Amount float64
}
