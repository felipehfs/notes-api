package models

import "time"

type Note struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	OwnerId     string    `json:"ownerId,omitempty" validate:"required"`
}
