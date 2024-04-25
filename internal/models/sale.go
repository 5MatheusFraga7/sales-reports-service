package models

import "time"

type Sale struct {
	Id       int
	SelledAt time.Time
	Value    float64
	Product  string
}
