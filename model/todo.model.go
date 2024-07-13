package model

import "time"

type Todo struct {
	ID          string    `json:"id,omitempty" bson:"_id,omitempty"`
	Description string    `json:"description,omitempty" bson:"description,omitempty"`
	Favourite   bool      `json:"favourite,omitempty" bson:"favourite,omitempty"`
	Completed   bool      `json:"completed,omitempty" bson:"completed,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" bson:"updated_at,omitempty"`
}
