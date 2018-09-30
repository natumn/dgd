package main

import (
	"os"
)

func main() {
	app := New()

	os.Exit(
		app.Run(),
	)
}
