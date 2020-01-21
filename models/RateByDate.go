package models

type RateByDate struct {
	Time string            `json:"date"`
	Rate map[string]string `json:"rate"`
}
