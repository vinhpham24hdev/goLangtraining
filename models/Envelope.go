package models

type Envelope struct {
	Name map[string]AvgRate `json:"rate"`
}
