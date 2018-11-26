package testa

import "fmt"

// datastore-gen-doc: true
type A struct {
	Test1 string `json:"test1" datastore:"test1"`
	Test2 string `json:"test2" datastore:"-"`
	Test3 string `json:"test3" datastore:"test3,test3option"`
}

func main() {
	a := A{
		Test1: "test1string",
		Test2: "test2string",
		Test3: "test3string",
	}

	fmt.Printf("%+v", a)
}
