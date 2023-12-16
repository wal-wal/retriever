package entity

import "time"

type Mantra struct {
	Content   string
	Writer    string
	Id        int
	CreatedAt time.Time
}
