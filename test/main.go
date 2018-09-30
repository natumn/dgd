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
	desc := `
鬼の国の女王、竜胆 尊じゃ。

ばあちゃるらいばあとして皆を楽しませられるように尽力する所存じゃ
わらわにそちらの世界のことをたくさん教えてくれ
`

	profile := NewProfile(user.ID, desc)

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", profile)
}
