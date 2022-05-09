package models

import "time"

type Expense struct {
	Id    int       `json:"Id"`
	Type  string    `json:"Type"`
	Title string    `json:"Title"`
	Date  time.Time `json:"Date"`
	Rate  float32   `json:"Rate"`
}
