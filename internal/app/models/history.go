package models

type History struct {
	ID     string  `json:"id,omitempty"`
	Time   string  `json:"time"`
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
