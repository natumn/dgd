package main

import (
	"fmt"
	"time"
)

type DatastoreUser struct {
	ID        int64     `datastore:"-" json:"id"`
	Name      string    `datastore:"name" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

func main() {
	user := &DatastoreUser{
		ID:        12345,
		Name:      "Rindou Mikoto",
		CreatedAt: time.Now(),
	}

	fmt.Printf("%+v\n", user)
}
