package entity

import (
	"github.com/google/uuid"
	"time"
)

type Mantra struct {
	Content   string
	Writer    string
	Id        uuid.UUID
	CreatedAt time.Time
}
