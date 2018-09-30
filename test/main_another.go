package main

import "time"

type Profile struct {
	UserID      int64     `json:"user_id" datastore:"user_id"`
	Description string    `json:"description" datastore:"description"`
	CreatedAt   time.Time `json:"created_at" datastore:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" datastore:"updated_at"`
}

func NewProfile(id int64, desc string) *Profile {
	return &Profile{
		UserID:      id,
		Description: desc,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
